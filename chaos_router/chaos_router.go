package chaos_router

import (
	"go_chaos/chaos_rule"
	"go_chaos/http_util"
	"go_chaos/routing_rules"
	"go_chaos/tcp_client"
	"go_chaos/tcp_server"
	"math/rand"
	"net"
	"time"
)

type ChaosRouter struct {
	port             int
	tcpServer        tcp_server.Server
	routingRules     []routing_rules.RouteRule
	chaosRules       []chaos_rule.ChaosRule
	handlers         []tcp_server.TcpHandler
	handlerLength    int
	maxPossibilities int
	errHandler       tcp_server.ErrHandler
}

func NewChaosRoute(
	port int,
	maxPossibilities int,
	routingRules []routing_rules.RouteRule,
	chaosRules []chaos_rule.ChaosRule,
	errHandler tcp_server.ErrHandler,
) tcp_server.Server {
	rand.Seed(time.Now().UnixNano())
	return &ChaosRouter{
		port:             port,
		maxPossibilities: maxPossibilities,
		routingRules:     routingRules,
		chaosRules:       chaosRules,
		errHandler:       errHandler,
	}
}

func (c ChaosRouter) Start() {
	
	// this is kind of messy
	var tcpHandler tcp_server.TcpHandler = c
	server := tcp_server.NewTcpServer(c.port, &tcpHandler, &c.errHandler)
	server.Start()
	c.tcpServer = server
}

func (c ChaosRouter) Stop() {
	c.tcpServer.Stop()
}

// tcp handler
func (c ChaosRouter) Handle(con net.Conn, errHandler *tcp_server.ErrHandler) {
	defer con.Close()
	c.initIfNeeded()
	randomInt := rand.Intn(c.handlerLength)
	if hdl := c.handlers[randomInt]; hdl == nil {
		// routeRequest
		tcpRequestString := tcp_server.ReadTcpRequestAsString(con)
		httpRequest := http_util.NewLazyHttpRequest(tcpRequestString)
		route := FindRouteRule(httpRequest, c.routingRules)
		if route == nil {
			// todo add better support for things we don't know how to handle
			return
		}
		
		// send request to routed
		response, err := tcp_client.WriteContentsAndGetResponseAsString(route.GetRoute().Host, route.GetRoute().Port, tcpRequestString)
		if err != nil {
			(*errHandler)(err)
		}
		
		// write back to original requester
		if _, err := con.Write([]byte(response)); err != nil {
			(*errHandler)(err)
		}
	} else {
		hdl.Handle(con, errHandler)
	}
}

func (c *ChaosRouter) initIfNeeded() {
	if c.handlers == nil {
		
		// todo validate array bounds maybe
		handlersByPercent := make([]tcp_server.TcpHandler, c.maxPossibilities)
		offset := 0
		for i := range c.chaosRules {
			currentRule := c.chaosRules[i]
			FillTcpHandlerArray(offset, currentRule.Percentage(), handlersByPercent, currentRule)
		}
		c.handlers = handlersByPercent
		c.handlerLength = len(c.handlers)
	}
}

func FindRouteRule(request http_util.HttpRequest, rules []routing_rules.RouteRule) routing_rules.RouteRule {
	for i := range rules {
		if rules[i].Applies(request) {
			return rules[i]
		}
	}
	
	// todo better default
	return nil
}

func FillTcpHandlerArray(start, end int, dest []tcp_server.TcpHandler, source tcp_server.TcpHandler) {
	for i := start; i < end; i++ {
		dest[i] = source
	}
}
