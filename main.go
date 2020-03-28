package main

import (
	"go_chaos/tcp_server"
	"net"
)

func main() {
	tcpServer := tcp_server.NewTcpServer(9000, func(con net.Conn, errHandler tcp_server.ErrHandler) {
		buffer := make([]byte, 2048)
		_, err := con.Read(buffer)
		if err != nil {
			errHandler(err)
		}
		println(string(buffer))
	},
		func(err error) {
			if err == nil {
				return
			}
			println(err.Error())
		},
	)
	tcpServer.Start()
}
