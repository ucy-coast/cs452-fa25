
package webserver

import (
	"encoding/json"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"sort"
	"sync"
	"strings"
	"unicode"
	"github.com/ucy-coast/websearch/internal/rpc_api"
)

// RootHandler serves the static HTML file
func (s *WebServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile(s.htmlPath)
	if err != nil {
		http.Error(w, "Unable to load HTML file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

// SearchHandler queries all shard RPC servers concurrently and merges the results
func (s *WebServer) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}

	results, err := s.searchAcrossShards(query, 10)
	if err != nil {
		http.Error(w, "search failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// searchAcrossShards sends the search request to all shard servers and aggregates the results.
func (s *WebServer) searchAcrossShards(query string, topK int) ([]rpc_api.SearchResult, error) {
	// Query all shards
	args := rpc_api.SearchArgs{
		Query: query,
		TopK:  topK,
	}

	var wg sync.WaitGroup
	mu := sync.Mutex{}
	var shardResults [][]rpc_api.SearchResult

	for _, addr := range s.shardRPCAddrs {
		wg.Add(1)
		go func(address string) {
			defer wg.Done()

			client, err := rpc.Dial("tcp", address)
			if err != nil {
				log.Printf("RPC dial failed to %s: %v", address, err)
				return
			}
			defer client.Close()

			var reply []rpc_api.SearchResult
			err = client.Call("IndexServer.Search", &args, &reply)
			if err != nil {
				log.Printf("RPC call failed to %s: %v", address, err)
				return
			}

			mu.Lock()
			shardResults = append(shardResults, reply)
			mu.Unlock()
		}(addr)
	}

	wg.Wait()

	finalResults := aggregateShardResults(shardResults, topK)
	
	return finalResults, nil
}
// aggregateShardResults merges results from multiple shards.
// If a document appears in multiple shards, its keyword matches are combined.
// Results are sorted by match count (descending), then document name (ascending).
func aggregateShardResults(shardResults [][]rpc_api.SearchResult, k int) []rpc_api.SearchResult {
	docMatches := make(map[string]map[string]struct{})

	// Merge keyword matches per document across all shards
	for _, results := range shardResults {
		for _, res := range results {
			if _, ok := docMatches[res.Document]; !ok {
				docMatches[res.Document] = make(map[string]struct{})
			}
			for _, kw := range res.Matches {
				docMatches[res.Document][kw] = struct{}{}
			}
		}
	}

	// Collect document names and sort for deterministic iteration
	docs := make([]string, 0, len(docMatches))
	for doc := range docMatches {
		docs = append(docs, doc)
	}
	sort.Strings(docs)

	// Build the merged and cleaned result list
	allResults := make([]rpc_api.SearchResult, 0, len(docMatches))
	for _, doc := range docs {
		kwSet := docMatches[doc]

		kwList := make([]string, 0, len(kwSet))
		for kw := range kwSet {
			kwList = append(kwList, kw)
		}
		sort.Strings(kwList)

		cleanedDoc := strings.TrimFunc(doc, func(r rune) bool {
			return unicode.IsSpace(r) || r == '\u00A0'
		})

		allResults = append(allResults, rpc_api.SearchResult{
			Document: cleanedDoc,
			Matches:  kwList,
		})
	}

	// Stable sort by descending number of matches, then by document name
	sort.SliceStable(allResults, func(i, j int) bool {
		if len(allResults[i].Matches) == len(allResults[j].Matches) {
			return allResults[i].Document < allResults[j].Document
		}
		return len(allResults[i].Matches) > len(allResults[j].Matches)
	})

	// Cap to top K
	if k < 0 {
		k = 0
	}
	if k > len(allResults) {
		k = len(allResults)
	}

	return allResults[:k]
}


