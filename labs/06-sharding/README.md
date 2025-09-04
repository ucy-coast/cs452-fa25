# Lab: Distributed Web Search with RPC 

In this lab tutorial, you will transform your standalone web search server into a distributed search system composed of a frontend server and multiple backend shards, communicating via Go’s `net/rpc`. This modular design will enable parallel processing and lay the foundation for deployment on a cluster in the following lab.

## Prerequisites

### Setting up the Experiment Environment in Cloudlab

For this tutorial, you will be using a CloudLab profile that comes with the latest version of Go. 

Start a new experiment on CloudLab using the `multi-node-cluster` profile in the `UCY-COAST-TEACH` project, configured with a single physical machine node. 

Open a remote SSH terminal session to `node0`.

Verify that the profile has a working installation of Go by typing the following command:

```
$ go version
```

Confirm that the command prints the installed version of Go. If you don't have Go installed then just follow the [download and install](https://go.dev/doc/install) steps.

## Architecture Overview

You will break up your web search server into:

#### Frontend (HTTP server):

Serves the UI and accepts user queries.

Forwards queries to all backend shards via RPC.

Merges responses and returns the top-K results.

#### Backends (RPC servers):

Each loads a separate portion of the inverted index.

Handles search queries via an RPC Search method.

Returns a list of matching documents with match count.

<figure> <p align="center"><img src="assets/images/websearch-distributed.png" width="30%"></p> <figcaption><p align="center">Figure. Distributed architecture using Go RPC</p></figcaption> </figure>
This change lets the system scale across multiple nodes and perform searches in parallel.

## Project Setup

Navigate to the starter directory for this lab:

```bash
cd labs/06-websearch-rpc/starter
```

Familiarize yourself with the structure:

```bash
.
├── cmd/
│   ├── indexserver/      # Backend RPC server
│   └── webserver/        # Web frontend + RPC client
├── internal/
│   ├── indexserver/      # Backend RPC server logic
│   ├── invertedindex/    # Inverted index logic
│   ├── rpc_api/          # Shared structs for RPC
│   └── webserver/        # Web frontend + RPC client logic
├── test/                 # Testing scripts and index data
└── web/
    └── static/           # HTML UI
```

We’ve provided scaffolding to help you get started. You can run and test your components independently.

## Part 1: Define the RPC API

Open and inspect the file: `internal/rpc_api/search.go`.

In this file, we've already provided the definitions for:

- The request type: `SearchArgs`, which includes the search query and desired number of top results.

- The response type: `SearchResult`, which contains a document name and the list of matched keywords.

These types define the structure of communication between the frontend server and shard servers during a distributed search.

Understand how the provided data structures are used in the RPC-based search system:

- SearchArgs is sent to each shard to request relevant results.

- Each shard returns a list of SearchResult structs based on its local index.

- The frontend server collects and aggregates results from all shards before returning them to the user.

Things to think about:

- What does a search request need to contain for the shard to process it?

- Why does each result include a document and matched keywords, rather than just the document?

- How should the frontend combine overlapping results from multiple shards?

## Part 2: Backend Index Server

Implement the backend in `internal/indexserver`.

The backend server should:

1. Load an index file from a `-index` flag.
1. Listen on a specified address using `-rpc_addr`.
1. Register a type implementing the Search method.
1. Handle incoming Search requests over RPC.

> 💡 Use Go’s net/rpc standard library. See the example used in class as reference.

You can test a backend by starting it and connecting with rpc.Dial.

### Part 3: Frontend Web Server

Implement the frontend in `internal/webserver`.

It should:

1. Serve the static HTML page.
1. Accept search queries via /api/search?q=....
1. Connect to multiple backend servers using addresses passed via a `-shards` flag (comma-separated).
1. Dispatch the query to all backends in parallel.
1. Merge the results and return the top-K documents as JSON.

You will need to:

- Parse flags: address, topK, shard addresses, HTML path.
- Use goroutines and channels to make parallel RPC calls.
- Implement a merging function to compute top-K documents from all responses.

## Testing

For testing, you will simulate a distributed multi-node system locally by running multiple processes on the same machine, which simplifies debugging and makes testing easier.

This lab is designed to be fully testable locally. In future extensions or deployments, the same architecture can scale to run across real nodes in a cluster, such as with Kubernetes.

To support local testing, we provide all necessary resources to run and validate your distributed search system using pre-partitioned index shards.

- While a tool `tools/partition_index` is available if you want to partition a large inverted index yourself, for this lab please use the provided shard files.

- Bash scripts located in the `test/` directory automate launching multiple backend servers and the frontend, executing queries, and verifying the results.

#### Steps to test your implementation locally:

1. Use the provided index shard files found in the test/shards/ directory.
1. Launch two or more backend server processes on different ports, each loading a distinct shard.
1. Start the frontend server, configured to communicate with all backend servers.
1. Access the web UI in your browser or use curl to send search queries and confirm that the results from all backends are correctly merged.