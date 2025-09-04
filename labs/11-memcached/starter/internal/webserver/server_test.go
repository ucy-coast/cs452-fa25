
package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer_HTTPRequests(t *testing.T) {
	topK := 10

	// Create server instance

	s := &Server{
		htmlPath: "../../web/static/index.html",
		topK:     topK,
	}

	// Set up custom mux for isolated testing
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.RootHandler)
	mux.HandleFunc("/api/search", s.SearchHandler)

	// Create test server
	ts := httptest.NewServer(mux)
	defer ts.Close()

	t.Run("GET /", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("GET /api/search?q=won", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/api/search?q=won")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
