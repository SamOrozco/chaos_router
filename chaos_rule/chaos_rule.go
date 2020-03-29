package chaos_rule

import "go_chaos/tcp_server"

type ChaosRule interface {
	Percentage() int
	tcp_server.TcpHandler
}
