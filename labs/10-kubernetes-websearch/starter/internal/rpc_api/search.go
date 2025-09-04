
package rpc_api

// SearchArgs contains the parameters for a search request
type SearchArgs struct {
    Query string
    TopK  int
}

// SearchResult represents a single search result
type SearchResult struct {
    Document string
    Matches []string  // keywords matched for this doc on that shard
}
