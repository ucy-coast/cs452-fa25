# Lab: Container Orchestration with Kubernetes

This tutorial provides a walkthrough of the basics of the Kubernetes cluster orchestration system.

## Background

With modern web services, users expect applications to be available 24/7, and developers expect to deploy new versions of those applications several times a day. Containerization helps package software to serve these goals, enabling applications to be released and updated without downtime. 

[Kubernetes](https://kubernetes.io/) (also known as k8s or “kube”) is a production-ready, open source container orchestration platform that automates many of the manual processes involved in deploying, managing, and scaling containerized applications. Kubernetes is designed with Google's accumulated experience in container orchestration, combined with best-of-breed ideas from the community.

### Kubernetes Cluster Architecture  

Kubernetes coordinates a highly available cluster of computers that are connected to work as a single unit. Kubernetes allows you to deploy containerized applications to a cluster without tying them specifically to individual machines. Kubernetes automates the distribution and scheduling of application containers across a cluster in a more efficient way.

<figure>
  <p align="center"><img src="assets/images/k8s-arch3.png" width="60%"></p>
  <figcaption><p align="center">Figure. Kubernetes Architecture</p></figcaption>
</figure>

A Kubernetes cluster is divided into two components:
- *Control plane*: coordinates the cluster.
- *Nodes*: run your application workloads.

#### Control plane

The [control plane](https://kubernetes.io/docs/concepts/overview/components/#control-plane-components) is responsible for managing the cluster. The control plane coordinates all activities in your cluster, such as scheduling applications, maintaining applications' desired state, scaling applications, and rolling out new updates.

The control plane includes the following core Kubernetes components:

- *kube-apiserver*:	The API server is how the underlying Kubernetes APIs are exposed. This component provides the interaction for management tools, such as kubectl or the Kubernetes dashboard.
- *kube-controller-manager*: The Controller Manager oversees a number of smaller controllers that perform actions such as replicating pods and handling node operations.
- *etcd*:	To maintain the state of a Kubernetes cluster and configuration, the highly available etcd is a key value store within Kubernetes.
- *kube-scheduler*:	When you create or scale applications, the Scheduler determines what nodes can run the workload and starts them.

Control plane components can be run on any machine in the cluster. However, for simplicity, set up scripts typically start all control plane components on the same machine, and do not run user containers on this machine. See [Creating Highly Available clusters with kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/high-availability/) for an example control plane setup that runs across multiple machines.

#### Nodes

A Kubernetes cluster consists of a set of worker machines, called [nodes](https://kubernetes.io/docs/concepts/architecture/nodes/), that run containerized applications. A worker node may be either a virtual or a physical machine, depending on the cluster. Each worker node is managed by the control plane. Every cluster has at least one worker node.

Each node runs the following components:

- *kubelet*: The Kubernetes agent that processes the orchestration requests from the control plane along with scheduling and running the requested containers.
- *kube-proxy*: Handles virtual networking on each node. The proxy routes network traffic and manages IP addressing for services and pods.
- *container runtime*: Allows containerized applications to run and interact with additional resources, such as the virtual network and storage. Docker is the default container runtime.

#### Putting it all together

The control plane manages the cluster and the nodes that are used to host the running applications.

When you deploy applications on Kubernetes, you tell the control plane to start the application containers. The control plane schedules the containers to run on the cluster's nodes. The nodes communicate with the control plane using the Kubernetes API, which the control plane exposes. End users can also use the Kubernetes API directly to interact with the cluster.

### Kunernetes Abstractions

Kubernetes uses different abstractions to represent the state of the system, such as pods, deployments, services, namespaces, and volumes.

#### Pods

A pod is a collection of containers sharing a network, acting as the basic unit of deployment in Kubernetes. All containers in a pod are scheduled on the same node.

Kubernetes uses pods to run an instance of your application. A pod represents a single instance of your application.

Pods typically have a 1:1 mapping with a container. In advanced scenarios, a pod may contain multiple containers. Multi-container pods are scheduled together on the same node, and allow containers to share related resources.

<figure>
  <p align="center"><img src="assets/images/kubernetes-pod.png" width="40%"></p>
  <figcaption><p align="center">Figure. Kubernetes Pod</p></figcaption>
</figure>


#### Deployments

A deployment is a supervisor for pods, giving you fine-grained control over how and when a new pod version is rolled out as well as rolled back to a previous state.  

#### Services

A service is an abstraction for pods, providing a stable, so called virtual IP (VIP) address. While pods may come and go and with it their IP addresses, a service allows clients to reliably connect to the containers running in the pod using the VIP. The "virtual" in VIP means it is not an actual IP address connected to a network interface, but its purpose is purely to forward traffic to one or more pods. Keeping the mapping between the VIP and the pods up-to-date is the job of kube-proxy, a process that runs on every node, which queries the API server to learn about new services in the cluster.

<figure>
  <p align="center"><img src="assets/images/kubernetes-service.png" width="40%"></p>
  <figcaption><p align="center">Figure. Kubernetes Service</p></figcaption>
</figure>

#### Namespaces

Namespaces provide a scope for Kubernetes resources, carving up your cluster in smaller units. You can think of it as a workspace you're sharing with other users. Many resources such as pods and services are namespaced. Others, such as nodes, are not namespaced, but are instead treated as cluster-wide. As a developer, you'll usually use an assigned namespace, however admins may wish to manage them, for example to set up access control or resource quotas.

#### Volumes

A Kubernetes volume is essentially a directory accessible to all containers running in a pod. In contrast to the container-local filesystem, the data in volumes is preserved across container restarts. The medium backing a volume and its contents are determined by the volume type:  

## Prerequisites

Before starting this tutorial, you will need access to a Kubernetes cluster. There are three common ways to get one set up:

#### Option 1: Use Minikube to Run Kubernetes Locally

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

#### Option 2: Use an Existing Kubernetes Cluster on CloudLab

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

#### Option 3: Create a New Kubernetes Cluster on CloudLab

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

## First contact with kubectl

To interact with Kubernetes during this tutorial we’ll use the Kubernetes command-line tool, `kubectl`. Kubectl communicates with a Kubernetes cluster's control plane using the Kubernetes API. We can use kubectl to deploy applications, inspect and manage cluster resources, and view logs. 

The most common operations can be done with the following kubectl commands:

- **kubectl get** - list resources
- **kubectl describe** - show detailed information about a resource
- **kubectl logs** - print the logs from a container in a pod
- **kubectl exec** - execute a command on a container in a pod
- **kubectl apply** -  apply a configuration to a resource by filename or stdin

### kubectl get nodes

To view the nodes in the cluster, run the ```kubectl get nodes``` command:

```
kubectl get nodes
```

This command shows all nodes that can be used to host our applications. Now we have only one node, and we can see that its status is ready (it is ready to accept applications for deployment).


## Running our first containers on Kubernetes

First things first: we cannot run a container. We are going to run a pod, and in that pod there will be a single container. In that container in the pod, we are going to run an Nginx webserver.

### Starting a simple pod

We can use `kubectl run` to start a single pod. We need to specify at least a *name* and the *image* we want to use. Optionally, we can specify the command to run in the pod.

Let’s run a pod that launches an Nginx webserver:

```bash
kubectl run webserver --image=nginx
```

The output tells us that a Pod was created:

```
pod/webserver created
```

Anything that the application would normally send to STDOUT becomes logs for the container within the Pod. We can retrieve these logs using the `kubectl logs` command.

Let's use the `kubectl logs` command to view our container's output. It takes a pod name as argument. Unless specified otherwise, it will only show logs of the first container in the pod.

View the output of our Nginx webserver:

```
kubectl logs webserver
```

Using standalone pods is generally frowned upon except for quick testing, as there is nothing in place to monitor the health of the pod.

### Scaling our application

We can use `kubectl scale` to scale a workload. The command takes the type of resource and the desired number of replicas as arguments: `kubectl scale TYPE NAME --replicas=HOWMANY`

Let's try it on our Pod, so that we have more Pods!

```
kubectl scale pod webserver --replicas=3
```

Alas, we get the following cryptic error:

```
Error from server (NotFound): the server could not find the requested resource
```

What's the meaning of that error? When we execute `kubectl scale THAT-RESOURCE --replicas=THAT-MANY`, it is like telling Kubernetes: go to `THAT-RESOURCE` and set the scaling button to position `THAT-MANY`. However, pods do not have a "scaling button". If we try to execute the `kubectl scale pod` command with `-v6`, we will see a PATCH request to /scale: that's the "scaling button". Technically it's called a subresource of the Pod.

As we see, we cannot "scale a Pod", although that's not completely true; we could give it more CPU/RAM. If we want more Pods, we need to create more Pods, that is execute `kubectl run` multiple times. There must be a better way!

There is a better way, indeed! We will create a ReplicaSet; a set of replicas is a set of identical pods. In fact, we will create a Deployment, which itself will create a ReplicaSet

Let's create a Deployment instead of a single Pod. 

```
kubectl create deployment webserver --image=nginx
```

Let's check the resources that were created. When you run `kubectl get all`, you will notice that a total of three new objects have been created. 

```
kubectl get all
```

```
NAME                            READY   STATUS    RESTARTS   AGE
pod/webserver                   1/1     Running   0          33s
pod/webserver-7c4f9bf7bf-mz4nr  1/1     Running   0          7s

NAME                        READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/webserver   1/1     1            1           8s

NAME                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/webserver-fc7cf666d   1         1         1       8s
```

We have the following resources:
- `deployment.apps/webserver`. This is the Deployment that we just created.
- `replicaset.apps/webserver-xxxxxxxxxx`. This is a Replica Set created by this Deployment.
- `pod/webserver-xxxxxxxxxx-yyyyy`. This is a pod created by the Replica Set.

When you created a Deployment, Kubernetes created a [Pod](https://kubernetes.io/docs/concepts/workloads/pods/) to host your application instance. A pod is a collection of containers sharing a network, acting as the basic unit of deployment in Kubernetes. All containers in a pod are scheduled on the same node. Our deployment controller manages a [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/), which it turn manages the pods, and ensures that the desired number are running.

Kubernetes Pods are mortal. Pods in fact have a lifecycle. When a worker node dies, the Pods running on the Node are also lost. A ReplicaSet might then dynamically drive the cluster back to desired state via creation of new Pods to keep your application running. As another example, consider an image-processing backend with 3 replicas. Those replicas are exchangeable; the front-end system should not care about backend replicas or even if a Pod is lost and recreated. That said, each Pod in a Kubernetes cluster has a unique IP address, even Pods on the same Node, so there needs to be a way of automatically reconciling changes among Pods so that your applications continue to function.

<figure>
  <p align="center"><img src="assets/images/module_02_first_app.svg" height="500"></p>
  <figcaption><p align="center">Figure. Deploying your first app on Kubernetes</p></figcaption>
</figure>

Let's try kubectl scale again, but on the Deployment. Scale our `webserver` deployment:

```
kubectl scale deployment webserver --replicas 3
```

Check that we now have multiple pods:

```
kubectl get pods
```

### Resilience

The deployment `webserver` watches its replica set. The replica set ensures that the right number of pods are running. What happens if pods disappear?

In a separate window, watch the list of pods:
  
```
watch kubectl get pods
```

Destroy the pod currently shown by kubectl logs:

```
kubectl delete pod webserver-xxxxxxxxxx-yyyyy
```

The command `kubectl delete pod` terminates the pod gracefully, meaning that it sends to the pod the TERM signal and waits for the pod to shutdown. 

As soon as the pod is in "Terminating" state, the Replica Set replaces it. But we can still see the output of the "Terminating" pod in `kubectl logs` until 30 seconds later, when the grace period expires. The pod is then killed, and `kubectl logs` exits

What happens if we delete a standalone pod, like the first `webserver` pod that we created? Let's find out, delete the pod:

```
kubectl delete pod webserver
```

We find that no replacement Pod gets created because there is no controller watching it. That's why we will rarely use standalone Pods in practice.

### Labels and selectors

[Labels and selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels) is a grouping primitive that allows logical operation on objects in Kubernetes. 

Labels are key/value pairs attached to objects and can be used in any number of ways:
- Designate objects for development, test, and production
- Embed version tags
- Classify an object using tags

Labels are arbitrary strings, with some limitations. The key must start and end with a letter or digit, can also have `.` `-` `_` (but not in first or last position), and can be up to 63 characters, or 253 + `/` + 63. The label *value* is up to 63 characters, with the same restrictions. 

Labels can be attached to objects at creation time or later on. They can be modified at any time. 

When we created the `webserver` Deployment, the deployment creation generated automatically a label for our Deployment and related Pods. 

With `kubectl describe deployment` command, you can see the label for our deployment:

```
kubectl describe deployment webserver
```

We see one label, `Labels: app=webserver`, which is added by `kubectl create deployment`. 

With the `kubectl describe pod` command, you can see the label for our pods:

```
kubectl describe pod webserver-xxxxxxxxxx-yyyyy
```

We see two labels:

```
Labels: app=webserver
        pod-template-hash=xxxxxxxxxx
```

The `app=webserver` label comes from `kubectl create deployment` too, while the `pod-template-hash` label was assigned by the Replica Set. 

We can use label selectors to identify a set of objects. A *selector* is an expression matching labels. It will restrict a command to the objects matching at least all these labels.

For example, we can list all the pods with at least `app=webserver`:

```
kubectl get pods --selector=app=webserver
```

We can also list all the pods with a label `app`, regardless of its value:

```
kubectl get pods --selector=app
```

You can do the same to list the existing deployments:

```
kubectl get deployments --selector=app=webserver
```

To apply a new label, we use the `kubectl label` command followed by the object type, object name and the new label:

```
kubectl label deployment webserver version=v1
```

This will apply a new label to our deployment (we pinned the application version to the Pod), and we can check it with the describe pod command:

```
kubectl describe deployment webserver
```

We see here that the label is attached now to our deployments. And we can query now the list of deployments using the new label:

```
kubectl get deployments --selector=version=v1
```

`kubectl get` gives us a couple of useful flags to check labels. `kubectl get --show-labels` shows all labels. `kubectl get -L xyz` shows the value of label `xyz`

List all the labels that we have on pods:

```
kubectl get pods --show-labels
```

List the value of label app on these pods:

```
kubectl get pods -L app
```

If a selector has multiple labels, it means "match at least these labels", for example: `--selector=app=frontend,release=prod`

The `--selector` flag can be abbreviated as `-l` (for labels):

```
kubectl get pods -l=app=webserver
```

We can also use negative selectors, for example: `--selector=app!=clock`

Selectors can be used with most `kubectl` commands. Examples: `kubectl delete`, `kubectl label`, ...

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

#### Exposing services for internal access

To create a new service and expose it to internal traffic we’ll use the `kubectl expose` command with the right port as parameter.

Let’s start with the most basic type of service: ClusterIP. This gives us an internal IP address reachable from within the cluster — not from the outside world.

Run this:

```bash
kubectl expose deployment webserver --port=8080 --target-port=80
```

Note: we didn’t specify `--type`, so Kubernetes used the default: `ClusterIP`.

Check the service:

```bash
kubectl get services
```

You’ll see something like this:

```
NAME        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE
webserver   ClusterIP   10.96.132.157    <none>        8080/TCP    5s
```

We now have an internal IP that load-balances across all three webserver pods.

Testing the internal service (with a curl pod)
To test this service, we need to make a request from inside the cluster. A common trick is to launch a temporary pod that includes curl or wget.

Let’s start an interactive pod using the curlimages/curl container:

```bash
kubectl run curl --image=curlimages/curl -it --restart=Never -- sh
```

Inside the container, run:

```sh
curl webserver:8080
```

You should see the HTML content of the Nginx welcome page.

#### Exposing services for external access

To create a new service and expose it to external traffic we’ll use the expose command with `NodePort` as parameter.

```bash
kubectl expose deployment webserver --port=8080 --target-port=80 --type=NodePort --name=webserver-nodeport
```

```
NAME                 TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
webserver            ClusterIP  10.96.132.157    <none>        8080/TCP         2m
webserver-nodeport   NodePort   10.96.142.201    <none>        8080:31234/TCP   10s
```

Take note of the NodePort (e.g. 31234).

Now you can access your service externally via:

```bash
curl http://<node-ip>:31234
```


