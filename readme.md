# Golang "Chaos" Router

**Go Chaos is a reverse proxy that can add little chaos(disruption) to your applications.**

Go chaos will answer the questions you've been asking about how your resilient you applications are in a distributed environment
prone to failures from external services. 

How to run your Go Chaos router binary
```./chaos_router C:\location\to\your\config\file```
You need one thing to get started a **configuration file**



## Configuration file 

Here is a simple configuration example.

The server will start on port 9002 and start listening for any http connections 
it has a maxPossibilities of 100, it has two routing rules and one chaos rule.

```json
{
  "port": 9002,
  "max_possibilities": 100,
  "routing_rules_configs": [
    {
      "match": {
        "type": "path",
        "match_value": "/my-ws/collection/",
        "match_type": "equals"
      },
      "route": {
        "host": "192.168.0.34",
        "port": 9001
      }
    },
    {
      "match": {
        "type": "path",
        "match_value": "that-ws",
        "match_type": "contains"
      },
      "route": {
        "host": "192.168.0.68",
        "port": 9002
      }
    }
  ],
  "chaos_rule_configs": [
    {
      "percent": 10,
      "response_status_code": 200,
      "ResponseBody": "success"
    }
  ]
}
```  
