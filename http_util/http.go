package http_util

type HttpRequest interface {
	Path() string
	Headers() map[string]string
}

type LazyHttpRequest struct {
	tcpRequest    string
	tcpHttpParser HttpParser
}

func NewLazyHttpRequest(tcpRequest string) HttpRequest {
	return &LazyHttpRequest{
		tcpRequest:    tcpRequest,
		tcpHttpParser: NewTcpHttpParser(),
	}
}

func (l LazyHttpRequest) Path() string {
	return l.tcpHttpParser.ParsePath(l.tcpRequest)
}

func (l LazyHttpRequest) Headers() map[string]string {
	return l.tcpHttpParser.ParseHeaders(l.tcpRequest)
}
