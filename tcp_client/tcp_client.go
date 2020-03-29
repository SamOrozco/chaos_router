package tcp_client

import (
	"fmt"
	"go_chaos/tcp_server"
	"net"
)

func WriteContentsAndGetResponseAsString(host string, port int, contents string) (string, error) {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return "", err
	}
	if _, err := con.Write([]byte(contents)); err != nil {
		return "", err
	}
	return tcp_server.ReadTcpRequestAsString(con), nil
}
