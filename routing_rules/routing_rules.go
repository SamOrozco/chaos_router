package routing_rules

import "go_chaos/http_util"

type Route struct {
	Host string
	Port int
}

type RouteRule interface {
	Applies(request http_util.HttpRequest) bool
	GetRoute() Route
}
