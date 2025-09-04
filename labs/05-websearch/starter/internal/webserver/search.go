
package webserver

import (
	"net/http"
)

// RootHandler serves HTML content loaded from a file
func (s *WebServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement me
}

// SearchHandler handles GET requests to the `/api/search` endpoint.
// It expects a query parameter `?q=...` containing space-separated keywords.
func (s *WebServer) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement me
}
