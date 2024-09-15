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

type TCPTransportOpts struct {
	listenAddress string
	handshaker    HandshakerFunc
	deocder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
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
		go t.HandleConn(conn)

		fmt.Printf("handling incomming connection %s\n", conn)
	}
}

type Temp struct{}

func (t *TCPTransport) HandleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.handshaker(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	msg := &Temp{}

	for {
		if err := t.deocder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error :%s\n", err)
			continue
		}
	}

}
