
package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	// TODO: implement shard search and result aggregation
	return nil, fmt.Errorf("searchAcrossShards: not implemented")
}


