
package webserver

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

// RootHandler serves HTML content loaded from a file
func (s *WebServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile(s.htmlPath)
	if err != nil {
		http.Error(w, "Unable to load HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

// SearchHandler handles GET requests to the `/api/search` endpoint.
// It expects a query parameter `?q=...` containing space-separated keywords.
func (s *WebServer) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}

	keywords := strings.Fields(strings.ToLower(query))
	results, err := s.index.SearchInMemoryTopK(keywords, s.topK)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
