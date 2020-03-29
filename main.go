package main

import (
	"go_chaos/chaos_router"
	"go_chaos/chaos_rule"
	"go_chaos/routing_rules"
	"go_chaos/util"
)

func main() {
	//
	// 	server := util.NewTestTcpServer(9002, `HTTP/1.1 200 OK
	// Date: Sun, 18 Oct 2009 08:56:53 GMT
	// Server: Apache/2.2.14 (Win32)
	// Last-Modified: Sat, 20 Nov 2004 07:16:26 GMT
	// ETag: "10000000565a5-2c-3e94b66c2e680"
	// Accept-Ranges: bytes
	// Content-Length: 44
	// Connection: close
	// Content-Type: text/html
	// X-Pad: avoid browser bug
	//
	// <html><body><h1> port 9002</h1></body></html>`)
	//
	// 	server.Start()
	
	// fake path rules
	pathRule := routing_rules.NewPathRule(util.Equal, "/athena/test", routing_rules.Route{
		Host: "192.168.0.34",
		Port: 9001,
	})
	
	pathRule1 := routing_rules.NewPathRule(util.Equal, "/hermes/test", routing_rules.Route{
		Host: "192.168.0.34",
		Port: 9002,
	})
	
	// fake chaos rule
	chaosRule := chaos_rule.BadRequestRule{}
	chaosRouter := chaos_router.NewChaosRoute(9000, 100, []routing_rules.RouteRule{pathRule, pathRule1}, []chaos_rule.ChaosRule{chaosRule}, func(err error) {
		println(err.Error())
	})
	chaosRouter.Start()
}
