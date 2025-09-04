
package indexserver

import (
	"fmt"

	"github.com/ucy-coast/websearch/internal/rpc_api"
)

func (s *IndexServer) Search(args *rpc_api.SearchArgs, reply *[]rpc_api.SearchResult) error {
	// TODO: implement search logic

	// Hint:
	// - Extract the query string from args.Query and split it into words (e.g., using strings.Fields).
	// - Normalize the words to lowercase for consistent matching.
	// - Call s.index.SearchInMemoryTopK(queryWords, args.TopK) to get top-K internal results.
	// - Convert each internal result into an rpc_api.SearchResult.
	// - Set *reply to the resulting slice.

	return fmt.Errorf("Search: not implemented")
}
