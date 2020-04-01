package tcp_server

import "net"

type DebugTcpHandler struct {
	response string
}

func NewDebugTcpHandler(response string) TcpHandler {
	return &DebugTcpHandler{response: response}
}

func (d DebugTcpHandler) Handle(con net.Conn, errHandler *ErrHandler) {
	println(ReadTcpRequestAsString(con))
	con.Write([]byte(d.response))
}

var DebugErrHandler ErrHandler = func(err error) {
	if err == nil {
		return
	}
	println(err.Error())
}

func ReadTcpRequestAsString(con net.Conn) string {
	buffer := make([]byte, 2048)
	if _, err := con.Read(buffer); err != nil {
		return ""
	}
	return string(buffer)
}

func ReadTcpRequestAsStringAsync(con net.Conn) chan string {
	responseChan := make(chan string, 0)
	go func() {
		buffer := make([]byte, 2048)
		if _, err := con.Read(buffer); err != nil {
			responseChan <- ""
		} else {
			responseChan <- string(buffer)
		}
	}()
	return responseChan
}
