---
title       : Deploying Web Search with Kubernetes
author      : Haris Volos
description : This is a hands-on look at deploying web search on a Kubernetes cluster.
keywords    : docker, containers, kubernetes
marp        : true
paginate    : true
theme       : jobs
--- 

<style>

  .img-overlay-wrap {
    position: relative;
    display: inline-block; /* <= shrinks container to image size */
    transition: transform 150ms ease-in-out;
  }

  .img-overlay-wrap img { /* <= optional, for responsiveness */
    display: block;
    max-width: 100%;
    height: auto;
  }

  .img-overlay-wrap svg {
    position: absolute;
    top: 0;
    left: 0;
  }

  </style>

  <style>
  img[alt~="center"] {
    display: block;
    margin: 0 auto;
  }

</style>

<style>   

   .cite-author {     
      text-align        : right; 
   }
   .cite-author:after {
      color             : orangered;
      font-size         : 125%;
      /* font-style        : italic; */
      font-weight       : bold;
      font-family       : Cambria, Cochin, Georgia, Times, 'Times New Roman', serif; 
      padding-right     : 130px;
   }
   .cite-author[data-text]:after {
      content           : " - "attr(data-text) " - ";      
   }

   .cite-author p {
      padding-bottom : 40px
   }

</style>


<style>
.code-small pre code {
  font-size: 0.7em;
}
</style>

<!-- _class: titlepage -->s: titlepage -->

# Lab: Deploying Web Search with Kubernetes

---

# Part 1: Deploying a Minimal Web Search Service

---

# Overview

- Deploy containerized apps to a Kubernetes cluster
- Use Kubernetes **Deployments** and **Services**
- Self-healing, resilient architecture
- Expose internal and external endpoints

---

# What is a Deployment?

- Describes how to **create and update** app instances
- Managed by the **Deployment Controller**
- Ensures high availability and **self-healing**
- Example: `kubectl create deployment ...`

---

# Deploying the Indexserver

Step 1: Generate a base manifest we can customize.

```bash
kubectl create deployment indexserver --image=hvolos01/indexserver \
  --dry-run=client -o yaml > indexserver-deployment.yaml
```

Step 2: Edit `indexserver-deployment.yaml` to add container args and port:

```yaml
args:
  - "--rpc_addr=:9090"
  - "--index_files=/index/invertedindex-medium.txt"
ports:
  - containerPort: 9090
```

---

# Deploying the Indexserver (cont'd)

Step 3: Launch the pod using our updated config:

```bash
kubectl apply -f indexserver-deployment.yaml
```

Step 4: Confirm deployment is running

```bash
kubectl get deployments
kubectl get pods
```

Should see the indexserver pod running

---

# Deploying the Webserver

Step 1: Generate a base manifest we can customize.

```bash
kubectl create deployment webserver --image=hvolos01/websearch_webserver \
  --dry-run=client -o yaml > webserver-deployment.yaml
```

Step 2: Edit `webserver-deployment.yaml` to add container args and port:

```yaml
args:
  - "--addr=0.0.0.0:8080"
  - "--shards=indexserver-service:9090"
  - "--topk=10"
  - "--htmlPath=index.html"
ports:
  - containerPort: 9090
```

---

# Deploying the Webserver (cont'd)

Step 3: Launch the pod using our updated config:

```bash
kubectl apply -f webserver-deployment.yaml
```

Step 4: Confirm deployment is running

```bash
kubectl get deployments
kubectl get pods
```

Should see the webserver pod running

---

# What is a Service?

- Abstracts access to a set of Pods

- Supports load balancing and discovery

- Types:

  - `ClusterIP` (default, internal-only)

  - `NodePort` (external access via node IP)
  - `LoadBalancer` (external IP, cloud support)

---

# Create Indexserver Service (internal access)

```bash
kubectl expose deployment indexserver \
  --name=indexserver-service \
  --type=ClusterIP \
  --port=9090 \
  --target-port=9090 \
```

---

# Expose Webserver (external access)

```bash
kubectl expose deployment webserver \
  --name=webserver-service \
  --type=NodePort \
  --port=80 \
  --target-port=8080
```

---

# Accessing the Web UI

Get the public hostname of the first node
```bash
PUBLIC_HOSTNAME=$(kubectl get nodes -o wide -o json | \
  jq -r '.items[0].status.addresses[] | select(.type=="Hostname") | .address')
export PUBLIC_HOSTNAME
```

Get the NodePort of the webserver-service
```bash
NODE_PORT=$(kubectl get svc webserver-service -o json | \
  jq -r '.spec.ports[0].nodePort')
export NODE_PORT
```

Output the full URL
```bash
echo "http://${PUBLIC_HOSTNAME}:${NODE_PORT}/"
```

---

# Cleaning up resources

Delete services and deployments

```bash
kubectl delete service webserver-service
kubectl delete service indexserver-service
kubectl delete deployment webserver
kubectl delete deployment indexserver
```

Verify:

```bash
kubectl get services
kubectl get deployments
```

---

# Part 2: Scaling and Sharding the Web Search Service

---

# Overview

- Sharding: Split index across multiple indexserver pods

- Replication: Run multiple pods for fault tolerance and scalable query performance

---

# Sharding the Indexserver - Shard 0

Generate and edit YAML manifest:

```bash
kubectl create deployment indexserver-shard-0 --image=hvolos01/websearch_indexserver \
  --dry-run=client -o yaml > indexserver-shard-0.yaml
```

```yaml
args:
  - "--rpc_addr=:9090"
  - "--index_files=/index/invertedindex-medium-0.txt"
ports:
  - containerPort: 9090
```
Create deployment and expose internal service:

```bash
kubectl apply -f indexserver-shard-0.yaml
kubectl expose deployment indexserver-shard-0 --port=9090 --target-port=9090 \
  --name=indexserver-shard-0-service --type=ClusterIP
```

---

# Sharding the Indexserver - Shard 1

Generate and edit YAML manifest:

```bash
kubectl create deployment indexserver-shard-1 --image=hvolos01/websearch_indexserver \
  --dry-run=client -o yaml > indexserver-shard-1.yaml
```

```yaml
args:
  - "--rpc_addr=:9090"
  - "--index_files=/index/invertedindex-medium-1.txt"
ports:
  - containerPort: 9090
```
Create deployment and expose internal service:

```
kubectl apply -f indexserver-shard-1.yaml
kubectl expose deployment indexserver-shard-1 --port=9090 --target-port=9090 \
  --name=indexserver-shard-1-service --type=ClusterIP
```

---

# Replicating the Shards

```bash
kubectl scale deployment indexserver-shard-0 --replicas=2
kubectl scale deployment indexserver-shard-1 --replicas=2

kubectl get pods -l app=indexserver-shard-0
kubectl get pods -l app=indexserver-shard-1
```

---

# Update Webserver for Sharded Backend

Edit `webserver-deployment.yaml`:

```yaml
args:
  - "--shards=indexserver-shard-0-service:9090,indexserver-shard-1-service:9090"
```

```bash
kubectl apply -f webserver-deployment.yaml
```

---

# Expose and Scale Webserver

```bash
kubectl expose deployment webserver \
  --type=NodePort --port=80 --target-port=8080 \
  --name=webserver-service

kubectl scale deployment webserver --replicas=3
kubectl get pods -l app=webserver
```

---

# Final Deployment Test

```bash
kubectl get deployments
kubectl get services
kubectl get pods
```

Visit your web UI via:

```bash
echo "http://${PUBLIC_HOSTNAME}:${NODE_PORT}/"
```
