# Golang "Chaos" Router

**Go Chaos is a reverse proxy that can add little chaos(disruption) to your applications.**

Go chaos will answer the questions you've been asking about how your resilient you applications are in a distributed environment
prone to failures from external services. 

How to run your Go Chaos router binary
```./chaos_router C:\location\to\your\config\file```
You need one thing to get started a **configuration file**



## Configuration file 

Here is a simple configuration example.

The server will start on port 9002 and start listening for any http connections, 
it has a maxPossibilities of 100, it has two routing rules and one chaos rule. All are explained in depth below.

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

### chaos_rule_configs

Chaos rule configurations are rules that will inject the specified response code and body by the given percent. 

```json
{
      "percent": 10,
      "response_status_code": 200,
      "ResponseBody": "success"
    }
```
**percent** = 10 times out of maxPossibilities return the specified status code and body

**response_status_code** = the status code of the response 

**http_response_body** = raw http response body 


This rule states 10 times out of 100(or max possibilities, it could have been anything) the router will return this response to any traffic that comes through it. 
*Rules might be expanded to be matched on a request similar to routing rules*



### routing rule configs

Routing rules match on a given path or header value. If no chaos rule is injected, the matched request will be routed to the specified route. The chaos router 
directly writes the routed servers response to the requester, as if nothing happened.

The routing request consists of two objects, the match and the route.
```json
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
    }
```

Route
```json
{
        "host": "192.168.0.34",
        "port": 9001
      }
```
**host** = host the request will be routed to

**port** = port to route the request to


Match

Path Matcher config
```json
{
        "type": "path",
        "match_value": "/my-ws/collection/",
        "match_type": "equals"
      }
```

Header Matcher Config
```json
{
        "type": "path",
        "match_value": "json_v2",
        "match_type": "equals", 
        "header_key": "Content-type"
      }
```
**type** = the type of matching you want to do (Path or header)

**match_value** = the value you're matching on

**match_type** = the type of matching you're doing (Equals or Contains, etc...)

**header_key**(only for header matching) = the header key you would like to match the value on
