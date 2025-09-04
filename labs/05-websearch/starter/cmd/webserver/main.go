
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ucy-coast/websearch/internal/invertedindex"
	"github.com/ucy-coast/websearch/internal/util"
	"github.com/ucy-coast/websearch/internal/webserver"
)

// parseFlags parses and validates command-line flags.
func parseFlags() (addr string, indexFiles []string, topK int, htmlPath string) {
	addrFlag := flag.String("addr", "0.0.0.0:8080", "IP address and port to listen on (e.g., :8080, 127.0.0.1:8080)")
	indexFilesFlag := flag.String("index", "", "Comma-separated list of index files to load")
	topKFlag := flag.Int("topk", 10, "Maximum number of top search results to return")
	htmlPathFlag := flag.String("htmlPath", "web/static/index.html", "Path to the HTML file to serve")
	flag.Parse()

	if *indexFilesFlag == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	validAddr, err := util.IsValidAddressWithDefaultPort(*addrFlag, 8080)
	if err != nil {
		log.Fatalf("Invalid address: %v", err)
	}

	files := strings.Split(*indexFilesFlag, ",")

	return validAddr, files, *topKFlag, *htmlPathFlag
}

func main() {
	addr, indexFiles, topK, htmlPath := parseFlags()

	index, err := invertedindex.NewInvertedIndex(indexFiles)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	s := webserver.NewWebServer(htmlPath, topK, index)

	http.HandleFunc("/", s.RootHandler)
	http.HandleFunc("/api/search", s.SearchHandler)

	fmt.Printf("Server running on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
