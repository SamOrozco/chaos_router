package tcp_server

import "net"

type DebugTcpHandler struct {
	response string
}

func (d DebugTcpHandler) Handle(con net.Conn, errHandler *ErrHandler) {
	println(ReadTcpRequestAsString(con))
	con.Write([]byte(d.response))
}

func NewDebugTcpHandler(response string) TcpHandler {
	return &DebugTcpHandler{response: response}
}

var DebugErrHandler ErrHandler = func(err error) {
	println(err.Error())
}

func ReadTcpRequestAsString(con net.Conn) string {
	buffer := make([]byte, 2048)
	if _, err := con.Read(buffer); err != nil {
		return ""
	}
	return string(buffer)
}
