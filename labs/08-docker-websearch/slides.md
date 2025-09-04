---
title       : Docker Containers
author      : Haris Volos
description : This is an introduction to Docker Containers
keywords    : docker, containers
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

<!-- _class: titlepage -->s: titlepage -->

# Lab: Deploying Web Applications with Docker

---

# Start a CloudLab experiment

- Go to your CloudLab dashboard
- Click on the Experiments tab
- Select Start Experiment
- Click on Change Profile
  - Select `multi-node-cluster` profile in the `UCY-COAST-TEACH` project
- Name your experiments with CloudLabLogin-ExperimentName
  - Prevents everyone from picking random names 

---

# Multi-Stage Docker Builds

A multi-stage Dockerfile uses multiple FROM instructions:

- Stage 1: Build the application

- Stage 2: Run the application

Benefits:

- Smaller images

- No dev tools in production

- Better security and performance

---

# Example: webserver.Dockerfile

## Stage 1: Build the application

```dockerfile
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .
RUN go build -o bin/webserver ./cmd/webserver
```

- Uses Go base image
- Copies source code
- Compiles the Go webserver binary to `/app/bin/webserver`

---

# Example: webserver.Dockerfile (cont'd)

## Stage 2: Run the application

```dockerfile
FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/webserver /webserver
COPY --from=builder /app/web/static/index.html /index.html

EXPOSE 8080
ENTRYPOINT ["/webserver"]
```

- Uses a minimal base (Ubuntu)
- Keeps final image small and clean

---

# Build the Image

Run the following:

```bash
docker build -t webserver -f webserver.Dockerfile .
```

Then verify:

```bash
docker images
```

---

# Running the Webserver Container 

Run on Default network:

```bash
docker run --name webserver -p 8888:8080 webserver \
  -addr :8080 \
  -htmlPath /index.html \
  -shards 127.0.0.1:9090
```

Access via:

```
http://<HOST_PUBLIC_ADDRESS>:8888
```

Cleanup:

```
docker rm webserver 
```

---

# Docker Bridge Networking and Port Mapping Explained

![w:800 center](figures/port-mapping.png)

- Webserver runs inside the container
- Listens on port 8080
- Host maps port 8888 → 8080

---

# Problem

Containers on default bridge network can’t talk to each other by name

Needed: A shared network where webserver can reach indexserver

---

# Create a User-Defined Network

Run the following:

```
docker network create websearch-net
```

Containers on network `websearch-net` can use each other’s names as hostnames

---

# Run Indexserver

```bash
docker run -d --name indexserver \
  --network websearch-net \
  -v $(pwd)/index:/index \
  indexserver \
  --rpc_addr=:9090 \
  --index_files=/index/invertedindex-medium.txt
```

- Mounts your host's `$(pwd)index/` folder
- Listens on port 9090

---

# Run Webserver (on same network)

```bash
docker run -d --name webserver -p 8888:8080 \
  --network websearch-net webserver \
  -addr :8080 \
  -htmlPath /index.html \
  -shards indexserver:9090
```

Now webserver can talk to indexserver using its name!

---

# Test It

In your browser:

```
http://<HOST_PUBLIC_ADDRESS>:8888
```

Try searching for a keyword like adventure

- Webserver sends query to the indexserver
- Indexserver processes the query and returns results to webserver
- Webserver formats and displays the results in the browser

---

# Cleanup

```bash
docker stop webserver indexserver
docker rm webserver indexserver
docker network rm websearch-net
```

---

# Docker Compose

---

# Why Docker Compose?

- One YAML file to define everything

- No need to run long docker run commands

- Simplifies multi-container setup

---

# docker-compose.yml

```yml
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
```

---

# docker-compose.yml (cont'd)

```yml
networks:
  webnet:
    driver: bridge
```

---

# Run the Services

To run:

```bash
docker compose up
```

- Builds the network

- Starts both services

- Connects everything automatically

To stop:

```bash
docker compose down
```

---

# Test Again

Open in your browser:

```
http://<HOST_PUBLIC_ADDRESS>:8888
```

Should behave just like before

All services connected via Compose

---

# Recap

- Built Docker images using multi-stage Dockerfiles

- Ran containers on isolated networks

- Connected services using Docker networks

- Simplified orchestration using Docker Compose

