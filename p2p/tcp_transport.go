package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}

}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshaker    HandShaker
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.AcceptLoop()

	return nil
}

func (t *TCPTransport) AcceptLoop() error {

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			return err
		}
		go t.handshaker.Handshake()
		go t.HandleConn(conn)
	}
}

func (t *TCPTransport) HandleConn(conn net.Conn) {
	fmt.Printf("handling incomming connection %s\n", t.listenAddress)
}

// func (t *TCPTransport) ReadMessage(conn net.Conn) {
// 	buf :=

// }
