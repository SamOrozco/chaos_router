package routing_rules

import (
	"go_chaos/http_util"
	"go_chaos/util"
	"testing"
)

type MockHttpRequest struct {
	path    string
	headers map[string]string
}

func NewMockHttpRequest(path string, headers map[string]string) http_util.HttpRequest {
	return &MockHttpRequest{
		path:    path,
		headers: headers,
	}
}

func (t MockHttpRequest) Path() string {
	return t.path
}

func (t MockHttpRequest) Headers() map[string]string {
	return t.headers
}

func TestNewPathEqualRuleApplies(t *testing.T) {
	
	// setup
	pathRule := NewPathRule(util.Equal, "/test/athena", Route{
		Host: "localhost",
		Port: 9000,
	})
	
	// execute
	rule1 := pathRule.Applies(NewMockHttpRequest("/test/", nil))
	rule2 := pathRule.Applies(NewMockHttpRequest("/test/athena", nil))
	rule3 := pathRule.Applies(NewMockHttpRequest("/tester/athena", nil))
	
	// verify
	if rule1 {
		t.Fail()
	}
	
	if !rule2 {
		t.Fail()
	}
	
	if rule3 {
		t.Fail()
	}
}

func TestNewHeaderEqualRuleApplies(t *testing.T) {
	// setup
	headerRule := NewHeaderRule("Connection", "keep", util.Equal, Route{
		Host: "localhost",
		Port: 8000,
	})
	
	rule1 := headerRule.Applies(NewMockHttpRequest("/no", map[string]string{
		"Connection": "keep-alive",
	}))
	
	rule2 := headerRule.Applies(NewMockHttpRequest("/no", map[string]string{
		"Connection": "keep",
	}))
	
	rule3 := headerRule.Applies(NewMockHttpRequest("/no", map[string]string{
		"Connection": "sleep",
	}))
	
	if rule1 {
		t.Fail()
	}
	
	if !rule2 {
		t.Fail()
	}
	
	if rule3 {
		t.Fail()
	}
	
}
