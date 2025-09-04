# Lab: Deploying Web Search with Kubernetes


## Prerequisites

Before starting this tutorial, you will need access to a Kubernetes cluster. There are three common ways to get one set up:

### Option 1: Use Minikube to Run Kubernetes Locally

If you want to run Kubernetes locally on your own laptop or workstation, Minikube is a popular choice.

1. Install Minikube following the official instructions:
   https://minikube.sigs.k8s.io/docs/start/

2. Start your local cluster:

    ```bash
    minikube start
    ```

3. Verify your cluster is running:

    ```bash
    kubectl version
    kubectl cluster-info
    ```
4. You can interact with Minikube services directly using:

    ```bash
    minikube service <service-name>
    ```

### Option 2: Use an Existing Kubernetes Cluster on CloudLab

If your instructor or lab administrator has already created a Kubernetes cluster for you on CloudLab:

1. Connect via SSH to the manager (control-plane) node of the cluster. This node’s hostname typically starts with kube1. Use the SSH command provided in the CloudLab interface. Example:

    ```bash
    ssh -p 22 alice@ms1019.utah.cloudlab.us
    ```
2. Once connected, verify `kubectl` is properly configured and talking to the cluster:

    ```bash
    kubectl version
    kubectl cluster-info
    ```

3. Check that your personal namespace exists:

    ```bash
    kubectl get namespaces | grep ${USER}
    ```

4. Confirm your `kubectl` context is set to your personal namespace:

    ```bash
    kubectl config view --minify | grep namespace:
    ```

### Option 3: Create a New Kubernetes Cluster on CloudLab

If you don’t have an existing cluster and want to create one:

1. Log in to CloudLab at `https://cloudlab.us`

2. Start a new experiment using the kubernetes profile under the `UCY-COAST-TEACH` project.

3. Wait for the cluster to fully initialize. This can take upwards of ten minutes.

4. Once ready, connect via SSH to the control-plane node (`kube1`...) using your CloudLab credentials:

    ```bash
    ssh -p 22 alice@ms1019.utah.cloudlab.us
    ```

5. Verify `kubectl` is installed and talking to the cluster:

    ```bash
    kubectl version
    kubectl cluster-info
    ```

6. Check for your personal namespace:

    ```bash
    kubectl get namespaces | grep ${USER}
    ```

7. If your namespace does not exist and you have admin rights, create it:

    ```bash
    kubectl create namespace <insert-namespace-name-here>
    ```

8. Set your kubectl context to use your namespace by default:

    ```bash
    kubectl config set-context --current --namespace=<insert-namespace-name-here>
    ```

If you don’t have admin access to create namespaces, please ask your instructor for assistance.

## Part 1: Deploying a Minimal Web Search Service 

Once you have a running Kubernetes cluster, you can deploy your containerized applications on top of it. To do so, you create a Kubernetes [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/ configuration. The Deployment instructs Kubernetes how to create and update instances of your application. Once you've created a Deployment, the Kubernetes control plane schedules the application instances included in that Deployment to run on individual Nodes in the cluster.

Once the application instances are created, a Kubernetes Deployment Controller continuously monitors those instances. If the Node hosting an instance goes down or is deleted, the Deployment controller replaces the instance with an instance on another Node in the cluster. This provides a self-healing mechanism to address machine failure or maintenance.

### Deploying containers

Let’s deploy our Web Search app on Kubernetes with the `kubectl create deployment` command. We need to provide the deployment name and app image location (include the full repository url for images hosted outside Docker hub).

In previous labs, our application was running on a single node. We didn't need to ship images because we built and ran on the same machine node. However, now we want to run on a cluster, so we need to have the same image on all the nodes of the cluster, and therefore we need to ship these images. The easiest way to ship container images is to use a container registry. For now, we will use pre-built images available from DockerHub. 

Let's start by deploying the `indexserver` as a single pod using `kubectl create`.

Run the following command to create a deployment for the `indexserver`:

```bash
kubectl create deployment indexserver --image=hvolos01/websearch_indexserver --dry-run=client -o yaml > indexserver-deployment.yaml
```

Next, edit `indexserver-deployment.yaml` to:

- Add the necessary container args for the server to function properly.

- Expose the container port that the indexserver listens on e.g. port 9090.

- (Optional) Add any volume mounts if your setup requires persistent storage. The images come with sample indexs in `/index` so you can avoid persistent volume by using those. 

```yaml
      - image: hvolos01/websearch_indexserver
        name: indexserver
        args:
          - "--rpc_addr=:9090"
          - "--index_files=/index/invertedindex-medium.txt"
        ports:
        - containerPort: 9090
```

Apply the deployment:

```bash
kubectl apply -f indexserver-deployment.yaml
```

Check the deployment status:

```bash
kubectl get deployments
kubectl get pods
```

Similarly, deploy the webserver using kubectl create:

```bash
kubectl create deployment webserver --image=hvolos01/websearch_webserver --dry-run=client -o yaml > webserver-deployment.yaml
```

Next, edit `webserver-deployment.yaml` to add the appropriate container args and expose the correct port.

```yaml
      - image: hvolos01/websearch_webserver
        name: webserver
        args:
          - "--addr=0.0.0.0:8080"
          - "--shards=indexserver-service:9090"
          - "--topk=10"
          - "--htmlPath=index.html"
        ports:
        - containerPort: 8080
```

`indexserver-service` is the name of the indexserver service which we will create next.

Apply the deployment:

```bash
kubectl apply -f webserver-deployment.yaml
```

Check the deployment status:

```bash
kubectl get deployments
kubectl get pods
```

### Exposing containers through services

Pods that are running inside Kubernetes are running on a private, isolated network. By default they are visible from other pods and services within the same kubernetes cluster, through their IP address. However, we then need to figure out a lot of things, including how to look up the IP address of the pod(s), how to connect from outside the cluster, how to load balance traffic, and how to handle pod failures. Kubernetes has a resource type named *Service*, which addresses all these questions!

A Service in Kubernetes is an abstraction which defines a logical set of Pods and a policy by which to access them. Services enable a loose coupling between dependent Pods. A Service is defined using YAML (preferred) or JSON, like all Kubernetes objects. The set of Pods targeted by a Service is usually determined by a LabelSelector (see below for why you might want a Service without including selector in the spec).

A Service routes traffic across a set of Pods. Services are the abstraction that allows pods to die and replicate in Kubernetes without impacting your application. Discovery and routing among dependent Pods (such as the frontend and backend components in an application) are handled by Kubernetes Services.

Although each Pod has a unique IP address, those IPs are not exposed outside the cluster without a Service. Services allow your applications to receive traffic. Services can be exposed in different ways by specifying a type in the ServiceSpec:

- ClusterIP (default) - Exposes the Service on an internal IP in the cluster. This type makes the Service only reachable from within the cluster.
- NodePort - Exposes the Service on the same port of each selected Node in the cluster using NAT. Makes a Service accessible from outside the cluster using \<NodeIP>:\<NodePort>. Superset of ClusterIP.
- LoadBalancer - Creates an external load balancer in the current cloud (if supported) and assigns a fixed, external IP to the Service. Superset of NodePort.
- ExternalName - Maps the Service to the contents of the externalName field (e.g. foo.bar.example.com), by returning a CNAME record with its value. No proxying of any kind is set up. This type requires v1.7 or higher of kube-dns, or CoreDNS version 0.0.8 or higher.

More information about the different types of Services can be found in the [Using Source IP](https://kubernetes.io/docs/tutorials/services/source-ip/) tutorial. Also see [Connecting Applications with Services](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/).

Additionally, note that there are some use cases with Services that involve not defining `selector` in the spec. A Service created without `selector` will also not create the corresponding Endpoints object. This allows users to manually map a Service to specific endpoints. Another possibility why there may be no selector is you are strictly using `type: ExternalName`.

#### Exposing index service for internal access

To create an index service and expose it to internal traffic we’ll use the `kubectl expose` command with the right port as parameter.

```bash
kubectl expose deployment indexserver \
  --port=9090 \
  --target-port=9090 \
  --name=indexserver-service \
  --type=ClusterIP
```

Explanation:
- `--port=9090`: the port the service will expose inside the cluster.
- `--target-port=9090`: the port the container actually listens on.
- `--name=indexserver-service`: gives the service a name (defaults to deployment name if omitted).
- `--type=ClusterIP`: internal-only access

#### Exposing services for external access

To create a web service and expose it to external traffic we’ll use the expose command with `NodePort` as parameter.

Expose the `webserver` service:

```bash
kubectl expose deployment webserver \
  --type=NodePort \
  --port=80 \
  --target-port=8080 \
  --name=webserver-service
```

What this does:
- `--port=80`: the port the service listens on inside the cluster
- `--target-port=8080`: the port your webserver container is exposing
- `--node-port=30080`: fixed external port on the node (you choose this)
- `--type=NodePort`: exposes the service externally on each node

> ⚠️ Make sure the port (e.g., 30080) is not already in use by another service. If it is, Kubernetes will return an error.

Let’s list the current Services from our cluster:

```bash
kubectl get services
```

```bash
NAME                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
indexserver-service   ClusterIP   10.109.183.151   <none>        9090/TCP       71s
webserver-service     NodePort    10.101.138.153   <none>        80:32002/TCP   4s
```

We have now two runnings services, all of which received a unique cluster-IP. We see that the frontend service received the external port 30373, in addition to a cluster-IP.

To find out what port was opened externally (by the NodePort option), we can also run the describe service command:

```bash
kubectl describe services/webserver-service
```

To connect to the service, we need the Public IP address of one of the worker nodes and the NodePort of the Service. You can use a bash processor called `jq` to parse JSON from command line.

```bash
export PUBLIC_HOSTNAME=$(kubectl get nodes -o wide -o json | jq -r '.items[0].status.addresses | .[] | select( .type=="Hostname" ) | .address ')
echo $PUBLIC_HOSTNAME

export NODE_PORT=$(kubectl get svc webserver-service --output json | jq -r '.spec.ports[0].nodePort' )
echo $NODE_PORT
```

We can now connect to the external IP address using the worker public hostname and allocated node port to view the web UI. Open the web UI in your browser at `http://${PUBLIC_HOSTNAME}:${NODE_PORT}/`

### Deleting a service

To delete Services you can use the delete service command.

```
kubectl delete service webserver-service
```

Confirm that the service is gone:

```
kubectl get services
```

This confirms that our webserver service was removed. To confirm that route is not exposed anymore you can curl the previously exposed IP and port:

```bash
curl ${PUBLIC_HOSTNAME}:${NODE_PORT}
```

This proves that the app is not reachable anymore from outside of the cluster. You can confirm that the app is still running with a curl inside the pod:

```bash
kubectl exec -ti frontend-xxxxxxxxxx-yyyyy -- curl localhost:8080
```

We see here that the application is up. This is because the Deployment is managing the application. To shut down the webserver, you would need to delete the Deployment as well.

```bash
kubectl delete deployment webserver
```

### Cleaning up resources

Once you're done experimenting with the single-instance indexserver and webserver, it's good practice to clean up your Kubernetes resources before moving on to the next part of the lab.

This will stop the pods managed by the deployments:

```bash
kubectl delete deployment indexserver
kubectl delete deployment webserver
```

You can verify that the deployments and pods are gone:

```bash
kubectl get deployments
kubectl get pods
```

This will remove the internal and external services you created:

```bash
kubectl delete service indexserver-service
kubectl delete service webserver-service
```

You can verify that the deployments, pods, and services are gone:

```bash
kubectl get services
```

You should see: 

```
No resources found in alice namespace.
```

(Replace alice with your actual namespace)


## Part 2: Scaling and Sharding the Web Search Service

Now that you’ve deployed a single indexserver and webserver, let’s scale out and shard the architecture to resemble a more realistic, distributed web search system.

We'll:
- Split the indexserver into multiple shards, each hosting a part of the index.

- Add replication to the index shards for fault tolerance and load balancing.

- Scale the webserver deployment to multiple replicas.

### Sharding the Indexserver

To support a larger inverted index, we will shard the index across multiple indexserver deployments.

Let’s say we want to split our index into 2 shards:

```bash
indexserver-shard-0
```

```bash
indexserver-shard-1
```

Each will serve a different portion of the index.

Create Deployment for Shard 0 (0-indexing)

```bash
kubectl create deployment indexserver-shard-0 --image=hvolos01/websearch_indexserver --dry-run=client -o yaml > indexserver-shard-0.yaml
```

Edit `indexserver-shard-0.yaml` to:

- Set the correct args for the shard (e.g., use /index/invertedindex-medium-0.txt)

- Expose port 9090

```yaml
      - image: hvolos01/websearch_indexserver
        name: indexserver
        args:
          - "--rpc_addr=:9090"
          - "--index_files=/index/invertedindex-medium-0.txt"        
        ports:
        - containerPort: 9090
```

Apply:

```bash
kubectl apply -f indexserver-shard-0.yaml
```

Expose as a service:

```bash
kubectl expose deployment indexserver-shard-0 \
  --port=9090 \
  --target-port=9090 \
  --name=indexserver-shard-0-service \
  --type=ClusterIP
```

Repeat for Shard 1

```bash
kubectl create deployment indexserver-shard-1 --image=hvolos01/websearch_indexserver --dry-run=client -o yaml > indexserver-shard-1.yaml
```

Edit and apply similarly, then expose:

```bash
kubectl expose deployment indexserver-shard-1 \
  --port=9090 \
  --target-port=9090 \
  --name=indexserver-shard-1-service \
  --type=ClusterIP
```

You now have two distinct index shards, each accessible via its own internal service.

### Replicating Index Shards

To add fault tolerance, you can run multiple replicas of each index shard. Kubernetes will automatically load balance traffic across the replicas.

Scale up:

```bash
kubectl scale deployment indexserver-shard-0 --replicas=2
kubectl scale deployment indexserver-shard-1 --replicas=2
```

Check that pods are running:

```bash
kubectl get pods -l app=indexserver-shard-0
kubectl get pods -l app=indexserver-shard-1
```

Kubernetes services will round-robin across pods behind each shard service.

### Updating the Webserver to Use Shards

Now update the webserver deployment so that each pod queries all index shards.

Edit `webserver-deployment.yaml` to provide multiple --shards= arguments, one for each index shard service:

```yaml
        args:
          - "--addr=0.0.0.0:8080"
          - "--shards=indexserver-shard-0-service:9090,indexserver-shard-1-service:9090"
          - "--topk=10"
          - "--htmlPath=index.html"
```

Apply the update:

```bash
kubectl apply -f webserver-deployment.yaml
```

Expose the webserver service:

```bash
kubectl expose deployment webserver \
  --type=NodePort \
  --port=80 \
  --target-port=8080 \
  --name=webserver-service
```

### Scaling the Webserver

Scaling the webserver improves throughput and resilience by distributing query traffic across multiple pods.

Then, use the kubectl scale command to increase replicas:

```bash
kubectl scale deployment webserver --replicas=3
```

Verify:

```bash
kubectl get pods -l app=webserver
```

You should now see 3 pods running for the webserver.

Note: The webserver must know about all index shards. We will update its configuration after deploying the shards.

Test Your Deployment

Check that everything is running:

```bash
kubectl get deployments
kubectl get services
kubectl get pods
```

Get the public IP and port of the webserver-service (as in the previous section), and access the web UI in your browser.

You now have:
- A scalable frontend with multiple webserver pods
- A sharded, replicated indexserver backend

