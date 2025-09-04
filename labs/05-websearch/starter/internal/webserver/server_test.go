

package webserver

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucy-coast/websearch/internal/invertedindex"
)

const testIndexContents = `adventure: 8 pg-being_ernest.txt, pg-dorian_gray.txt, pg-emma.txt, pg-frankenstein.txt, pg-les_miserables.txt, pg-moby_dick.txt, pg-sherlock_holmes.txt, pg-war_and_peace.txt
mystery: 9 pg-being_ernest.txt, pg-dorian_gray.txt, pg-dracula.txt, pg-grimm.txt, pg-huckleberry_finn.txt, pg-les_miserables.txt, pg-sherlock_holmes.txt, pg-tale_of_two_cities.txt, pg-ulysses.txt
friendship: 10 pg-being_ernest.txt, pg-emma.txt, pg-frankenstein.txt, pg-great_expectations.txt, pg-grimm.txt, pg-les_miserables.txt, pg-moby_dick.txt, pg-sherlock_holmes.txt, pg-tale_of_two_cities.txt, pg-war_and_peace.txt
journey: 7 pg-dorian_gray.txt, pg-dracula.txt, pg-frankenstein.txt, pg-grimm.txt, pg-les_miserables.txt, pg-tom_sawyer.txt, pg-ulysses.txt
danger: 6 pg-dracula.txt, pg-frankenstein.txt, pg-grimm.txt, pg-les_miserables.txt, pg-moby_dick.txt, pg-ulysses.txt
betrayal: 8 pg-being_ernest.txt, pg-dorian_gray.txt, pg-emma.txt, pg-great_expectations.txt, pg-grimm.txt, pg-les_miserables.txt, pg-tale_of_two_cities.txt, pg-war_and_peace.txt
courage: 5 pg-being_ernest.txt, pg-dracula.txt, pg-frankenstein.txt, pg-les_miserables.txt, pg-tom_sawyer.txt
hope: 9 pg-being_ernest.txt, pg-dorian_gray.txt, pg-emma.txt, pg-grimm.txt, pg-huckleberry_finn.txt, pg-les_miserables.txt, pg-moby_dick.txt, pg-tale_of_two_cities.txt, pg-war_and_peace.txt
destiny: 7 pg-dorian_gray.txt, pg-dracula.txt, pg-emma.txt, pg-grimm.txt, pg-les_miserables.txt, pg-sherlock_holmes.txt, pg-tom_sawyer.txt
legacy: 6 pg-being_ernest.txt, pg-emma.txt, pg-frankenstein.txt, pg-grimm.txt, pg-moby_dick.txt, pg-ulysses.txt`

func createTempIndexFile(t *testing.T, contents string) string {
	t.Helper() // Marks this function as a test helper

	tmpFile, err := os.CreateTemp("", "invertedindex-*.txt")
	require.NoError(t, err)

	_, err = tmpFile.WriteString(contents)
	require.NoError(t, err)

	err = tmpFile.Close()
	require.NoError(t, err)

	t.Cleanup(func() {
		os.Remove(tmpFile.Name())
	})

	return tmpFile.Name()
}

func TestServer_HTTPRequests(t *testing.T) {
	topK := 10
	indexFile := createTempIndexFile(t, testIndexContents)
	indexFiles := []string{indexFile}

	// Create server instance
	index, err := invertedindex.NewInvertedIndex(indexFiles)
	require.NoError(t, err)

	s := &Server{
		htmlPath: "../../web/static/index.html",
		index:    index,
		topK:     topK,
	}

	// Set up custom mux for isolated testing
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.rootHandler)
	mux.HandleFunc("/api/search", s.searchHandler)

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
