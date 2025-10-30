
package indexserver

import (
	"github.com/ucy-coast/websearch/internal/invertedindex"
)

// IndexServer exposes search functionality over RPC.
type IndexServer struct {
	index *invertedindex.InvertedIndex
}

func NewIndexServer(index *invertedindex.InvertedIndex) *IndexServer {
	s := &IndexServer{
		index:      index,
	}
	return s
}
