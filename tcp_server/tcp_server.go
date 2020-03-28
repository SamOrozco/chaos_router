package tcp_server

import (
	"fmt"
	"net"
)

type Server interface {
	Start()
	Stop()
}

type ErrHandler func(err error)
type TcpHandler func(con net.Conn, errHandler ErrHandler)

type TcpServer struct {
	port       int
	handler    TcpHandler
	errHandler ErrHandler
	stopChan   chan bool
	Resilient  bool
}

func NewTcpServer(port int, handler TcpHandler, errHandler ErrHandler) Server {
	return &TcpServer{
		port:       port,
		handler:    handler,
		errHandler: errHandler,
		stopChan:   make(chan bool, 0),
		Resilient:  true,
	}
}

func (t *TcpServer) Start() {
	t.stopChan = make(chan bool, 0)
	t.startAsyncTcpRequestHandler(t.port, t.errHandler)
	<-t.stopChan
}

func (t TcpServer) Stop() {
	t.stopChan <- true
}

func (t TcpServer) startAsyncTcpRequestHandler(port int, errHandler ErrHandler) {
	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		errHandler(err)
	}
	go t.handleTcpRequestAsync(tcpListener, errHandler)
}

func (t TcpServer) handleTcpRequestAsync(listener net.Listener, errHandler ErrHandler) {
	for {
		con, err := listener.Accept()
		if err != nil {
			errHandler(err)
			if !t.Resilient {
				t.Stop()
			}
		}
		go t.handlerWrapper(con, errHandler)
	}
}

func (t TcpServer) handlerWrapper(con net.Conn, handler ErrHandler) {
	t.handler(con, handler)
	handler(con.Close())
}
