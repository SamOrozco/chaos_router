package chaos_rule

import (
	"go_chaos/tcp_server"
	"net"
)

const badRequestText string = `HTTP/1.1 400 OK
<html><body>Bad Request</h1></body></html>`

type BadRequestRule struct {
}

func (b BadRequestRule) Percentage() int {
	return 5
}

func (b BadRequestRule) Handle(con net.Conn, errHandler *tcp_server.ErrHandler) {
	if _, err := con.Write([]byte(badRequestText)); err != nil {
		(*errHandler)(err)
	}
}
