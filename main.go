package main

import (
	config2 "go_chaos/config"
	"go_chaos/routing_rules"
)

func main() {
	config := config2.ChaosRouterConfig{
		Port:             9001,
		MaxPossibilities: 100,
		RoutingRulesConfigs: []config2.RoutingRuleConfig{
			{
				Match: config2.MatchConfig{
					Type:       "path",
					MatchValue: "/athena/test",
					MatchType:  "equals",
				},
				Route: routing_rules.Route{
					Host: "192.168.0.34",
					Port: 9001,
				},
			},
			{
				Match: config2.MatchConfig{
					Type:       "path",
					MatchValue: "/athena",
					MatchType:  "equals",
				},
				Route: routing_rules.Route{
					Host: "192.168.0.34",
					Port: 9002,
				},
			},
		},
		ChaosRuleConfigs: []config2.ChaosRuleConfig{
			{
				Percent:            50,
				ResponseStatusCode: 200,
				ResponseBody:       "success",
			},
			{
				Percent:            50,
				ResponseStatusCode: 400,
				ResponseBody:       "bad request",
			},
		},
	}
	
	chaosRouter := config2.CreateChaosRouterFromConfig(config)
	chaosRouter.Start()
}
