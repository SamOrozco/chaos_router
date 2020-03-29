package routing_rules

import (
	"go_chaos/http_util"
	"testing"
)

// setup
var TestHttpRequestString string = `GET /test/athena HTTP/1.1
Host: localhost:9000
Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36
Sec-Fetch-Dest: document
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-Site: cross-site
Sec-Fetch-Mode: navigate
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9`

func TestTcpHttpParser_ParsePath(t *testing.T) {
	parser := http_util.NewTcpHttpParser()
	if parser.ParsePath(TestHttpRequestString) != "/test/athena" {
		t.Fail()
	}
}

func TestTcpHttpParser_ParseHeaders(t *testing.T) {
	parser := http_util.NewTcpHttpParser()
	headers := parser.ParseHeaders(TestHttpRequestString)
	if len(headers) != 11 {
		t.Fail()
	}
	
	if headers["Connection"] != "keep-alive" {
		t.Fail()
	}
	
	if headers["Cache-Control"] != "max-age=0" {
		t.Fail()
	}
}
