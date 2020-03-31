package main

import (
	config2 "go_chaos/config"
	"go_chaos/routing_rules"
)

func main() {
	config := config2.ChaosRouterConfig{
		Port:             9001,
		MaxPossibilities: 1000,
		RoutingRulesConfigs: []config2.RoutingRuleConfig{{
			Match: config2.MatchConfig{
				Type:       "path",
				MatchValue: "/athena/test",
				MatchType:  "equals",
			},
			Route: routing_rules.Route{
				Host: "192.168.0.34",
				Port: 9001,
			},
		}},
		ChaosRuleConfigs: []config2.ChaosRuleConfig{{
			Percent:            500,
			ResponseStatusCode: 400,
			ResponseBody:       "bad request",
		}},
	}
	
	chaosRouter := config2.CreateChaosRouterFromConfig(config)
	chaosRouter.Start()
}
