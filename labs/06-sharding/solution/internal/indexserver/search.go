
package indexserver

import (
	"fmt"
	"strings"

	"github.com/ucy-coast/websearch/internal/rpc_api"
)

func (s *IndexServer) Search(args *rpc_api.SearchArgs, reply *[]rpc_api.SearchResult) error {
	queryWords := strings.Fields(strings.ToLower(args.Query))

	// Use the new index method to get top-k results with matched keywords
	results, err := s.index.SearchInMemoryTopK(queryWords, args.TopK)
    if s.useMemory {
        results, err = s.index.SearchInMemoryTopK(queryWords, args.TopK)
    } else {
        results, err = s.index.SearchScanFiles(queryWords)
    }
	if err != nil {
		return fmt.Errorf("search failed: %v", err)
	}

	// Convert internal Result to rpc_api.SearchResult
	var rpcResults []rpc_api.SearchResult
	for _, res := range results {
		rpcResults = append(rpcResults, rpc_api.SearchResult{
			Document: res.Document,
			Matches: res.Matches,
		})
	}

	*reply = rpcResults
	return nil
}
