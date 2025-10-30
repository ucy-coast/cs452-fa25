
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/ucy-coast/websearch/internal/util"
	"github.com/ucy-coast/websearch/internal/webserver"
)

// parseFlags parses and validates CLI flags
func parseFlags() (addr string, shardRPCAddrs []string, topK int, htmlPath string) {
	addrFlag := flag.String("addr", "0.0.0.0:8080", "Address to listen on (e.g., :8080, 127.0.0.1:8080)")
	shardsFlag := flag.String("shards", "", "Comma-separated list of shard RPC addresses (e.g., 127.0.0.1:9090,127.0.0.1:9091)")
	topKFlag := flag.Int("topk", 10, "Maximum number of top search results to return")
	htmlPathFlag := flag.String("htmlPath", "web/static/index.html", "Path to HTML file to serve")

	flag.Parse()

	if *shardsFlag == "" {
		fmt.Fprintf(os.Stderr, "Usage: %s -shards addr1,addr2,... [OPTIONS]\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	validAddr, err := util.IsValidAddressWithDefaultPort(*addrFlag, 8080)
	if err != nil {
		log.Fatalf("Invalid server address: %v", err)
	}

	shardAddrs := strings.Split(*shardsFlag, ",")
	for _, s := range shardAddrs {
		if _, _, err := net.SplitHostPort(s); err != nil {
			log.Fatalf("Invalid shard RPC address: %s", s)
		}
	}
	return validAddr, shardAddrs, *topKFlag, *htmlPathFlag
}

func main() {
	addr, shardRPCAddrs, topK, htmlPath := parseFlags()
	s := webserver.NewWebServer(htmlPath, topK, shardRPCAddrs)

	http.HandleFunc("/", s.RootHandler)
	http.HandleFunc("/api/search", s.SearchHandler)

	fmt.Printf("Webserver running on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}


