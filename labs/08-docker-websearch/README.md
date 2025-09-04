# Lab: Deploying Web Applications with Docker 

In this lab, you will deploy your web search application with Docker

## Prerequisites 

### Setting up the Experiment Environment in Cloudlab

For this tutorial, you can use a CloudLab image that already has the Docker Engine installed.

#### 1. Start a New CloudLab Experiment

Start a new experiment on CloudLab using the `multi-node-cluster` profile in the `UCY-COAST-TEACH` project, configured with a single physical machine node. 

#### 2. Retrieve the WebSearch Code

If you still have your code from previous labs, use that. Otherwise, you can clone the starter code using:

```bash
git clone https://github.com/ucy-coast/cs452-fa25.git
cd cs452-fa25/labs/08-docker-websearch/starter
```

⚠️ Important: Clone the repository on the physical machine where you’ll be running Docker commands. Do not clone it inside a Docker container.

## Building Docker Images with Multi-Stage Dockerfiles

In this section, we will learn how to build a Docker image for our websearch service using a multi-stage Dockerfile. This approach helps us create smaller, more efficient container images by separating the build environment from the runtime environment.

### What is the Multi-Stage Build Pattern?

Multi-stage builds allow us to use multiple `FROM` statements in a single Dockerfile. Each `FROM` defines a stage:
- The first stage contains all the tools and dependencies needed to build the application (like the Go compiler).
- The final stage contains only the runtime environment and the compiled application binary.

By copying only the necessary artifacts from the build stage into the final image, we avoid including the full build environment in the final image, which reduces image size, improves security by excluding build tools and source code, and simplifies deployment with only what's needed to run the app.

### Webserver

We begin by building the Docker image for the webserver. Below is the complete `Dockerfile` used to create this image. You can save it as `webserver.Dockerfile` in your project directory.

The next step now is to create an image with this web app. As mentioned above, all user images are based on a base image. Since our application is written in Go, the base image we're going to use will be [Go (golang)](https://hub.docker.com/_/golang).

The application directory does contain a Dockerfile but since we're doing this for the first time, we'll create one from scratch. To start, create a new blank file in our favorite text-editor and save it in the same folder as the flask app by the name of ``Dockerfile``.

```dockerfile
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o bin/webserver ./cmd/webserver

FROM ubuntu:22.04

# Install certificates and any other needed dependencies in one RUN layer
RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/webserver /webserver
COPY --from=builder /app/web/static/index.html /index.html

EXPOSE 8080

ENTRYPOINT ["/webserver"]
```

Here is a step-by-step explanation of our Dockerfile

#### Stage 1: Builder

```dockerfile
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o bin/webserver ./cmd/webserver
```

- Uses the official `golang:1.24` image which has Go installed.
- Sets the working directory inside the container to `/app`.
- Copies the entire project source code from the host into the container.
- Runs the go build command to compile the Go code located in `./cmd/webserver` and outputs the binary executable at `/app/bin/webserver`.

#### Stage 2: Final runtime image

```dockerfile
FROM ubuntu:22.04

# Install certificates and any other needed dependencies in one RUN layer
RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    rm -rf /var/lib/apt/lists/*
```

- Starts from a clean `ubuntu:22.04` base image, which is much smaller than the Go image.
- Installs only the necessary runtime dependencies: `curl` and SSL certificates (ca-certificates).
- Cleans up package manager caches to keep the image small.

```dockerfile
COPY --from=builder /app/bin/webserver /webserver
COPY --from=builder /app/web/static/index.html /index.html
```
- Copies the compiled Go binary from the builder stage into the root of the final image.
- Copies a static HTML file used by the webserver into the container.

```dockerfile
EXPOSE 8080

ENTRYPOINT ["/webserver"]
```

- `EXPOSE` 8080 tells Docker and other tools that the container listens on port `8080`.
- `ENTRYPOINT` sets the default executable to run when the container starts — here it runs the `/webserver` binary.

Now that we have our `Dockerfile`, we can build our image:

```bash
$ docker build -t websearch_webserver -f webserver.Dockerfile .
```

Sample output:
```
Step 1/10 : FROM golang:1.24 AS builder
1.24: Pulling from library/golang
ebed137c7c18: Pull complete 
c2e76af9483f: Pull complete 
37f838b71c6b: Pull complete 
a73486c29d94: Pull complete 
a3e4aa2eec44: Pull complete 
85db7d6e6763: Pull complete 
4f4fb700ef54: Pull complete 
Digest: sha256:ef5b4be1f94b36c90385abd9b6b4f201723ae28e71acacb76d00687333c17282
Status: Downloaded newer image for golang:1.24
 ---> f14dd5573539
Step 2/10 : WORKDIR /app
 ---> Running in 74613750566a
 ---> Removed intermediate container 74613750566a
 ---> 678ac5239876
Step 3/10 : COPY . .
 ---> 42891bcc97de
Step 4/10 : RUN go build -o bin/webserver ./cmd/webserver
 ---> Running in 1331bbf9b169
 ---> Removed intermediate container 1331bbf9b169
 ---> cd2689d120b2
Step 5/10 : FROM ubuntu:22.04
22.04: Pulling from library/ubuntu
1d387567261e: Pull complete 
Digest: sha256:1ec65b2719518e27d4d25f104d93f9fac60dc437f81452302406825c46fcc9cb
Status: Downloaded newer image for ubuntu:22.04
 ---> 1d3ca894b30c
Step 6/10 : RUN apt-get update &&     apt-get install -y curl ca-certificates &&     rm -rf /var/lib/apt/lists/*
 ---> Running in a90677271e6c
 ---> Removed intermediate container a90677271e6c
 ---> 24600f81f85f
Step 7/10 : COPY --from=builder /app/bin/webserver /webserver
 ---> d2cad9b8ba62
Step 8/10 : COPY --from=builder /app/web/static/index.html /index.html
 ---> d095bb0de5a0
Step 9/10 : EXPOSE 8080
 ---> Running in 0310c757bdc7
 ---> Removed intermediate container 0310c757bdc7
 ---> 218edb5a7502
Step 10/10 : ENTRYPOINT ["/webserver"]
 ---> Running in 6c1a3e15f0e6
 ---> Removed intermediate container 6c1a3e15f0e6
 ---> 1220fa7d0609
Successfully built 1220fa7d0609
Successfully tagged websearch_webserver:latest
```

If you don't have the ``golang:1.24`` image, the client will first pull the image and then create your image. Hence, your output from running the command will look different from mine. If everything went well, your image should be ready! Run ``docker images`` and see if your image shows.


### Indexserver

Repeat the same steps to build the Docker image for the indexserver. Create a separate `Dockerfile` for it, then build the image and name it indexserver.

## Running the Containers and Setting Up Networking

### Step 1: Run the Webserver Container on the Default Network

Before creating any custom networks, let's start by running the webserver container on Docker’s default bridge network. This simple setup helps you understand the basics of running containers.

Run the following command:

```bash
docker run --name webserver -p 8888:8080 websearch_webserver -addr :8080 -htmlPath /index.html -shards 127.0.0.1:9090
```

Let’s break this down:
- `--name webserver`: Assigns a custom name to the container, making it easier to reference.
- `-p 8888:8080`: Maps port **8080** inside the container (where the webserver listens) to port **8888** on your local machine (host).
- `websearch_webserver`: The name of the Docker image you previously built.
- The rest (`-addr :8080 -htmlPath /index.html -shards 127.0.0.1:9090`) are arguments passed to the webserver binary inside the container.

If all goes well, you should see a ``Webserver running on...`` message in your terminal. 

<figure>
  <p align="center"><img src="figures/port-mapping.png"></p>
  <figcaption><p align="center">Figure. Port mapping</p></figcaption>
</figure>

As shown above, the container has its own internal IP address. The webserver inside the container listens on port 8080 for HTTP requests. However, this internal port is not directly accessible from your machine (the Docker host).

That’s where the `-p` flag comes in. It creates a bridge between the host and container, mapping an external port (8888 in this case) on your machine to an internal port (8080) in the container.

So, by running this command, we’ve made the webserver inside the container accessible at:

```bash
http://<HOST_PUBLIC_ADDRESS>:8888
```

Head over to that URL in your browser, and you should see the web interface served by your Go-based webserver. 

When you're done experimenting, return to the terminal where the webserver is running and press `Ctrl+C` to stop the container.

This only stops the container; it doesn’t remove it. To clean up the stopped container, run:

```bash
docker rm webserver
```

### Step 2: Run Webserver and Indexserver on a Shared Network

When running multiple services (like a webserver and an indexserver), it’s important to enable container-to-container communication. Docker's default network doesn't allow containers to reach each other by name, which makes service coordination difficult.

To solve this, we’ll use a user-defined bridge network, which lets containers communicate using their names as hostnames.

#### Create a Docker Network

First, create a custom bridge network:

```bash
docker network create websearch-net
```

This creates a network named `websearch-net`.

Containers attached to this network can resolve each other by name, which is ideal for service discovery.

#### Run the Indexserver Container

The indexserver needs access to a an inverted index file. Let’s assume this file is located on your host machine at:

```
<your project>/index/invertedindex-medium.txt
```

To make this file available inside the container, we use Docker volumes to map a host directory into the container.

Run the following command:

```bash
docker run -d --name indexserver \
  --network websearch-net \
  -v $(pwd)/index:/index \
  websearch_indexserver \
  --rpc_addr=:9090 \
  --index_files=/index/invertedindex-medium.txt
```

Next, start the indexserver and attach it to the network:

Let’s break this down:
- `-d`: Runs the container in detached mode (in the background).
- `--name indexserver`: Names the container so other containers can refer to it.
- `--network websearch-net`: Connects the container to the shared network.
- `-v $(pwd)/index:/index`: Maps the `index` directory on your host to `/index` in the container. This gives the container access to the index file.
- `websearch_indexserver`: The image name you built earlier.
- The rest (`--rpc_addr=:9090 --index_files=/index/invertedindex-medium.txt`) are arguments passed to the webserver binary inside the container. Specifically, `--rpc_addr=:9090` tells the indexserver to listen for RPC requests on port 9090, and `--index_files=/index/invertedindex-medium.txt` points the indexserver to the mounted file inside the container.

This container doesn’t need a published port, since it's only accessed internally by the webserver.

#### Run the Webserver Container on the Same Network

Now, run the webserver container and also connect it to the same network:

```bash
docker run -d --name webserver -p 8888:8080 --network websearch-net webserver \
  -addr :8080 \
  -htmlPath /index.html \
  -shards indexserver:9090
```

This connects the webserver to the same websearch-net network. It passes the `-shards` argument pointing to `indexserver:8080`, using the container name as a hostname. Once started, the webserver will now be able to communicate with the indexserver at http://indexserver:8080.

#### Testing

Open your browser and navigate to the webserver URL. You should see the web interface served by your Go-based webserver. Try entering a keyword query like `adventure` and click Search. The webserver should now be able to communicate with the indexserver backend and return search results based on your query.

#### Cleanup

To stop and remove both containers:

```bash
docker stop webserver indexserver
docker rm webserver indexserver
```

And to remove the custom network:

```bash
docker network rm websearch-net
```

## Using Docker Compose to Run the Webserver and Indexserver

Instead of running multiple docker run commands and manually setting up networks and volumes, you can simplify the process using Docker Compose. Compose lets you define multi-container applications in a single YAML file and spin everything up with one command.

### Step 1: Create a `docker-compose.yml` file

Create a file named `docker-compose.yml` in your project directory and add the following:

```yaml
version: "3.8"

services:
  webserver:
    image: websearch_webserver
    container_name: webserver
    ports:
      - "8888:8080"
    command: -addr :8080 -htmlPath /index.html -shards indexserver:9090
    networks:
      - webnet
  
  indexserver:
    image: websearch_indexserver
    container_name: indexserver
    volumes:
      - ./index:/index
    command: --rpc_addr=:9090 --index_files=/index/invertedindex-medium.txt
    networks:
      - webnet

networks:
  webnet:
    driver: bridge
```

### Step 2: Run the Services

In the same directory as the `docker-compose.yml` file, run:

```bash
docker compose up
```

This will:
- Build the network webnet (if it doesn't already exist)
- Start both indexserver and webserver containers
- Map ports and volumes as needed
- Pass startup arguments via the command field

To stop everything:

```bash
docker compose down
```

### Step 3: Test

Open your browser and go to:

```
http://HOST_PUBLIC_ADDRESS:8888
```

Enter a keyword like adventure and click Search. The webserver will talk to the indexserver and return search results, just as before.
