package routing_rules

import "go_chaos/http_util"

type Route struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RouteRule interface {
	Applies(request http_util.HttpRequest) bool
	GetRoute() Route
	IsFallbackRoutingRule() bool
}
