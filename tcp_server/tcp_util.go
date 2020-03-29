package tcp_server

import "net"

func ReadTcpRequestAsString(con net.Conn) string {
	buffer := make([]byte, 2048)
	if _, err := con.Read(buffer); err != nil {
		return ""
	}
	return string(buffer)
}
