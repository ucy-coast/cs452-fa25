
package webserver

type WebServer struct {
	htmlPath      string
	topK          int
	shardRPCAddrs []string
}
func NewWebServer(htmlPath string, topK int, shardAddrs []string) *WebServer {
	s := &WebServer{
		htmlPath:      htmlPath,
		topK:          topK,
		shardRPCAddrs: shardAddrs,
	}
	return s
}


