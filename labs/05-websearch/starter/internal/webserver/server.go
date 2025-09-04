
package webserver

import (
	"github.com/ucy-coast/websearch/internal/invertedindex"
)

type WebServer struct {
	htmlPath string
	topK     int
	index    *invertedindex.InvertedIndex
}

func NewWebServer(htmlPath string, topK int, index *invertedindex.InvertedIndex) *WebServer {
	s := &WebServer{
		htmlPath: htmlPath,
		topK:     topK,
		index:    index,
	}
	return s
}
