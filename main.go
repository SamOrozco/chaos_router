package main

import (
	"go_chaos/chaos_router"
	"go_chaos/chaos_rule"
	"go_chaos/routing_rules"
	"go_chaos/tcp_server"
	"go_chaos/util"
)

func main() {
	go BuildListener()
	go BuildListener1()
	
	// fake path rules
	pathRule := routing_rules.NewPathRule(util.Equal, "/athena/test", routing_rules.Route{
		Host: "localhost",
		Port: 9001,
	})
	
	pathRule1 := routing_rules.NewPathRule(util.Equal, "/athena", routing_rules.Route{
		Host: "localhost",
		Port: 9002,
	})
	
	// fake chaos rule
	chaosRule := chaos_rule.BadRequestRule{}
	
	chaosRouter := chaos_router.NewChaosRoute(9000, 10, []routing_rules.RouteRule{pathRule, pathRule1}, []chaos_rule.ChaosRule{chaosRule}, func(err error) {
		println(err.Error())
	})
	chaosRouter.Start()
}

func BuildListener() {
	var tcpHandler = tcp_server.NewDebugTcpHandler(`HTTP/1.1 200 OK
Date: Sun, 18 Oct 2009 08:56:53 GMT
Server: Apache/2.2.14 (Win32)
Last-Modified: Sat, 20 Nov 2004 07:16:26 GMT
ETag: "10000000565a5-2c-3e94b66c2e680"
Accept-Ranges: bytes
Content-Length: 44
Connection: close
Content-Type: text/html
X-Pad: avoid browser bug

<html><body><h1>It works!</h1></body></html>`)
	tcpServer := tcp_server.NewTcpServer(9001, &tcpHandler, &tcp_server.DebugErrHandler)
	tcpServer.Start()
}

func BuildListener1() {
	var tcpHandler = tcp_server.NewDebugTcpHandler(`
HTTP/1.1 200 OK
Date: Sun, 18 Oct 2009 08:56:53 GMT
Server: Apache/2.2.14 (Win32)
Last-Modified: Sat, 20 Nov 2004 07:16:26 GMT
ETag: "10000000565a5-2c-3e94b66c2e680"
Accept-Ranges: bytes
Content-Length: 44
Connection: close
Content-Type: text/html
X-Pad: avoid browser bug

<html><body><h1>It kinda works!</h1></body></html>`)
	tcpServer := tcp_server.NewTcpServer(9002, &tcpHandler, &tcp_server.DebugErrHandler)
	tcpServer.Start()
}
