package routing_rules

import (
	"go_chaos/http_util"
	"go_chaos/util"
)

// PATH RULES
type PathRouteRule struct {
	matcher   util.StringMatcher
	ruleValue string
	Route     Route
}

func NewPathRule(
	matcher util.StringMatcher,
	ruleValue string,
	route Route,
) RouteRule {
	return &PathRouteRule{
		matcher:   matcher,
		ruleValue: ruleValue,
		Route:     route,
	}
}

func (p PathRouteRule) Applies(request http_util.HttpRequest) bool {
	return p.matcher(request.Path(), p.ruleValue)
}

func (p PathRouteRule) GetRoute() Route {
	return p.Route
}
