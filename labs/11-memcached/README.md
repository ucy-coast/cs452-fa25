# Lab: Memcached

This lab tutorial will introduce you to the Memcached memory-caching system. 

## Background

## Memcached

Memcached is a general-purpose distributed memory caching system. It is often used to speed up dynamic data-driven websites by caching data and objects in RAM to reduce the number of times an external data source (such as a database or API) must be read.

Memcached's APIs provide a very large hash table distributed across multiple machines. When the table is full, subsequent inserts cause older data to be purged in least recently used order. Applications using Memcached typically layer requests and additions into RAM before falling back on a slower backing store, such as a database.

## Caching Web Search Queries with Memcached

Search queries can be expensive because they often involve complex operations over large datasets. Caching popular queries reduces backend load and speeds up response times, improving user experience.

### Using Memcached as a Query Cache in Web Search

Memcached provides a simple set of operations (set, get, and delete) that makes it attractive as an elemental component in a large-scale distributed system. 

We will rely on Memcached to lighten the read load on our index servers. In particular, we will use Memcached as a demand-filled look-aside cache as shown in the Figure below. 

<figure>
  <p align="center"><img src="assets/images/memcached-look-aside-cache.png" width="80%"></p>
  <figcaption><p align="center">Figure. Memcached as a demand-filled look-aside cache. The figure illustrates the read path for a web server on a cache miss.</p></figcaption>
</figure>

Specifically, when a user submits a search query, the web application first checks Memcached using the query string as the key.

- If the data is present in the cache (cache hit), it is returned immediately, avoiding expensive backend queries.

- If the data is not cached (cache miss), the application fetches the data from the index service, returns it to the user, and stores it in Memcached for future requests.

When the search index or underlying data is updated, the corresponding cached entries should be invalidated or deleted to prevent stale results; however, this tutorial will not cover cache invalidation.

### Extend the Search Service to Use Memcached

As a first step, you will need to extend the `webserver` service to interact with Memcached for caching query results. 

Package [`memcache`](https://pkg.go.dev/github.com/bradfitz/gomemcache/memcache) provides a client library to connect to a Memcached server.

Add a command-line flag in your main server or CLI that receives the Memcached server address and pass this address when initializing the web service.

```go
mmcAddrFlag := flag.String("memc_addr", "localhost:11211", "Address of the memcached server")
```

Add a new field `cacheClient` to the `WebServer` type to hold the Memcached client instance:

```go
type WebServer struct {
	htmlPath      string
	topK          int
	shardRPCAddrs []string
	cacheClient   *memcache.Client
}
```

Second, extend the function `NewWebServer` to receive an additional parameter `mmcAddr` that corresponds to the address of the memcached service and initialize the client:

```go
memcache.New(mmcAddr)
```

In your search logic, that is in `searchAcrossShards`, query Memcached first for each search key:

```go
item, err := s.cacheClient.Get(cacheKey)
if err == nil {
    // Cache hit: unmarshal cached results
    var cachedResults []SearchResult
    if err := json.Unmarshal(item.Value, &cachedResults); err == nil {
        return cachedResults, nil
    }
} else if err != memcache.ErrCacheMiss {
    log.Warnf("Memcached error: %v", err)
}
```

If the cache miss occurs, proceed to retrieve the results from your index servers as usual.

After fetching results from the index servers, marshal the results and store them in Memcached:

```go
if jsonData, err := json.Marshal(finalResults); err == nil {
    err = s.cacheClient.Set(&memcache.Item{Key: cacheKey, Value: jsonData})
    if err != nil {
        log.Warnf("Failed to set cache for key %s: %v", cacheKey, err)
    }
}
```

You will also need to extend the import statement to include the additional packages  `github.com/bradfitz/gomemcache/memcache`.

### Memcached Service

With the `webserver` service updated to use Memcached, the next step is to set up a Memcached instance for it to connect to.
Depending on your deployment method, this service can be created using either Docker Compose or Kubernetes.

For example, with Docker Compose, you can extend the `docker-compose.yml` file in `deploy/docke`r` to add a new service named `memcached`:


```yaml
services:
  frontend:
  ...
  memcached:
    image: memcached
    container_name: 'websearch_memcached'
    restart: always
    environment:
      - MEMCACHED_CACHE_SIZE=128
      - MEMCACHED_THREADS=2
    networks:
      - websearchnet
```

The `memcached` service spins a Memcached server inside a container. 

The Memcached server uses 2 threads and sets aside `MEMCACHED_CACHE_SIZE` MB of RAM for the cache.

After building the webserver and indexserver images, you can run the app by invoking:

```
docker-compose up
```

To stop containers and remove containers, networks, volumes, and images created by `up`.

```
docker-compose down --volumes
```
