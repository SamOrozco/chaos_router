package http_util

import (
	"fmt"
	"net"
	"time"
)

const httpResponseString = `HTTP/1.1 %d %s
Now: %d
Content-Type: text/html

%s`

func WriteHttpResponseToTcpConnection(statusCode int, responseBody string, conn net.Conn) error {
	htmlString := fmt.Sprintf(
		httpResponseString,
		statusCode,
		getStatusStringFromStatusCode(statusCode),
		time.Now().Unix(),
		responseBody,
	)
	
	_, err := conn.Write([]byte(htmlString))
	if err != nil {
		return err
	}
	return nil
}

func getStatusStringFromStatusCode(statusCode int) string {
	return "OK"
}
