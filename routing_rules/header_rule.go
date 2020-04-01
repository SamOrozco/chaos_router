package routing_rules

import (
	"go_chaos/http_util"
	"go_chaos/util"
)

type HeaderRule struct {
	route     Route
	headerKey string
	headerVal string
	matcher   util.StringMatcher
}

func NewHeaderRule(
	headerKey,
	headerVal string,
	matcher util.StringMatcher,
	route Route,
) RouteRule {
	return &HeaderRule{
		route:     route,
		headerKey: headerKey,
		headerVal: headerVal,
		matcher:   matcher,
	}
}

func (h HeaderRule) Applies(request http_util.HttpRequest) bool {
	if val, exists := request.Headers()[h.headerKey]; exists {
		return h.matcher(val, h.headerVal)
	} else {
		return false
	}
}

func (h HeaderRule) GetRoute() Route {
	return h.route
}

func (h HeaderRule) IsFallbackRoutingRule() bool {
	return h.headerKey == "*" || h.headerVal == "*"
}
