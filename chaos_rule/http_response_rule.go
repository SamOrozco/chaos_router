package chaos_rule

import (
	"go_chaos/http_util"
	"go_chaos/tcp_server"
	"net"
)

type HttpResponseRule struct {
	statusCode   int
	responseBody string
	percent      int
}

func NewHttpResponseRule(statusCode int, responseBody string, percent int) ChaosRule {
	return &HttpResponseRule{
		statusCode:   statusCode,
		responseBody: responseBody,
		percent:      percent,
	}
}

func (h HttpResponseRule) Percentage() int {
	return h.percent
}

func (h HttpResponseRule) Handle(con net.Conn, errHandler *tcp_server.ErrHandler) {
	if err := http_util.WriteHttpResponseToTcpConnection(h.statusCode, h.responseBody, con); err != nil {
		(*errHandler)(err)
	}
}
