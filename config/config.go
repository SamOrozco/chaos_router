package config

import (
	"go_chaos/chaos_router"
	"go_chaos/chaos_rule"
	"go_chaos/routing_rules"
	"go_chaos/tcp_server"
	"go_chaos/util"
	"strings"
)

type ChaosRouterConfig struct {
	Port                int                 `json:"port"`
	MaxPossibilities    int                 `json:"max_possibilities"`
	RoutingRulesConfigs []RoutingRuleConfig `json:"routing_rules_configs"`
	ChaosRuleConfigs    []ChaosRuleConfig   `json:"chaos_rule_configs"`
}

/**
Match - how to match an http request
Route - where to match requests that match to
*/
type RoutingRuleConfig struct {
	Match MatchConfig         `json:"match"`
	Route routing_rules.Route `json:"route"`
}

/**
Type - [Path, or Header]
MatchValue - "/test/resource"
MatchType - Equals or contains
HeaderKey - if type is header this is the header key you are trying to match
*/
type MatchConfig struct {
	Type       string `json:"type"`
	MatchValue string `json:"match_value"`
	MatchType  string `json:"match_type"`
	HeaderKey  string `json:"header_key"`
}

/**
Percent - how often you want this rule called
ResponseStatusCode - response state code
ResponseBody - response body
ResponseHeaders - response headers
*/
type ChaosRuleConfig struct {
	Percent            int `json:"percent"`
	ResponseStatusCode int `json:"response_status_code"`
	ResponseBody       string
}

/**
Create a chaos router from the config format
*/
func CreateChaosRouterFromConfig(config ChaosRouterConfig) chaos_router.ChaosRouter {
	port := config.Port
	maxPoss := config.MaxPossibilities
	chaosRuleConfigs := config.ChaosRuleConfigs
	routeRuleConfig := config.RoutingRulesConfigs
	routingRules := getRouteRuleFromConfig(routeRuleConfig)
	chaosRules := getChaosRuleFromConfig(chaosRuleConfigs)
	return chaos_router.NewChaosRouter(port, maxPoss, routingRules, chaosRules, tcp_server.DebugErrHandler)
}

func getChaosRuleFromConfig(configs []ChaosRuleConfig) []chaos_rule.ChaosRule {
	if configs == nil || len(configs) == 0 {
		return nil
	}
	
	result := make([]chaos_rule.ChaosRule, 0)
	for _, config := range configs {
		percent := config.Percent
		responseBody := config.ResponseBody
		statusCode := config.ResponseStatusCode
		result = append(result, chaos_rule.NewHttpResponseRule(statusCode, responseBody, percent))
	}
	return result
}

func getRouteRuleFromConfig(configs []RoutingRuleConfig) []routing_rules.RouteRule {
	if configs == nil || len(configs) == 0 {
		return nil
	}
	
	result := make([]routing_rules.RouteRule, 0)
	
	// converting from config to real life
	for _, config := range configs {
		route := config.Route
		matchConfig := config.Match
		matchType := matchConfig.Type
		ruleValue := matchConfig.MatchValue
		stringMatcher := getStringMatcherFromType(matchConfig.MatchType)
		
		matchType = strings.ToLower(matchType)
		if matchType == "path" {
			result = append(result, routing_rules.NewPathRule(stringMatcher, ruleValue, route))
		}
	}
	return result
}

func getStringMatcherFromType(typeString string) util.StringMatcher {
	// todo this is not an extensible solution
	if typeString == "contains" || typeString == "cont" {
		return util.Contains
	}
	return util.Equal
}
