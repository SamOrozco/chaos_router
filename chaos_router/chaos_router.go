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
	port                int
	tcpServer           tcp_server.Server
	routingRules        []routing_rules.RouteRule
	chaosRules          []chaos_rule.ChaosRule
	handlers            []tcp_server.TcpHandler
	handlerLength       int
	maxPossibilities    int
	errHandler          tcp_server.ErrHandler
	fallbackRoutingRule routing_rules.RouteRule
	requestCleaner      func(string) string
}

func NewChaosRouter(
	port int,
	maxPossibilities int,
	routingRules []routing_rules.RouteRule,
	chaosRules []chaos_rule.ChaosRule,
	errHandler tcp_server.ErrHandler,
	requestCleaner func(val string) string,
) ChaosRouter {
	rand.Seed(time.Now().UnixNano())
	return ChaosRouter{
		port:             port,
		maxPossibilities: maxPossibilities,
		routingRules:     routingRules,
		chaosRules:       chaosRules,
		errHandler:       errHandler,
		requestCleaner: requestCleaner,
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
	
	// we have to read the tcp request before we can write
	// I am going to read from this request async
	requestStringChan := tcp_server.ReadTcpRequestAsStringAsync(con)
	defer close(requestStringChan)
	
	// alot of important login in this func
	c.initIfNeeded()
	randomInt := rand.Intn(c.handlerLength)
	if hdl := c.handlers[randomInt]; hdl == nil {
		// routeRequest
		// wait to finish writing
		tcpRequestString := <-requestStringChan
		httpRequest := http_util.NewLazyHttpRequest(tcpRequestString)
		route := FindRouteRule(httpRequest, c.routingRules)
		if route == nil && c.fallbackRoutingRule != nil {
			// fallback if possible
			route = c.fallbackRoutingRule
		} else if route == nil {
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
		// wait to finish writing
		<-requestStringChan
		hdl.Handle(con, errHandler)
	}
}

/**
clean request string if needed before sent to routed server
 */
func (c ChaosRouter) cleanRequest(request string) string {
	return c.requestCleaner(request)
}

/**
todo better documentation about the chaos rules
*/
func (c *ChaosRouter) initIfNeeded() {
	if c.handlers == nil {
		c.initChaosRules()
		c.initFallbackRoutingRuleIfExists()
	}
}

func (c *ChaosRouter) initChaosRules() {
	handlersByPercent := make([]tcp_server.TcpHandler, c.maxPossibilities)
	offset := 0
	for i := range c.chaosRules {
		currentRule := c.chaosRules[i]
		percent := currentRule.Percentage()
		FillTcpHandlerArray(offset, offset+percent, handlersByPercent, currentRule)
		offset = percent
	}
	c.handlers = handlersByPercent
	c.handlerLength = len(c.handlers)
}

func (c *ChaosRouter) initFallbackRoutingRuleIfExists() {
	if c.fallbackRoutingRule == nil {
		for _, v := range c.routingRules {
			if v.IsFallbackRoutingRule() {
				c.fallbackRoutingRule = v
				
				// there can only be one fallback routing rule
				return
			}
		}
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
