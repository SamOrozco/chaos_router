package util

import "go_chaos/tcp_server"

func NewTestTcpServer(port int, response string) tcp_server.Server {
	tcpHandler := tcp_server.NewDebugTcpHandler(response)
	return tcp_server.NewTcpServer(port, &tcpHandler, &tcp_server.DebugErrHandler)
}
