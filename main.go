package main

import (
	"log"

	"github.com/VAtulKumar1/Distributed-File-Storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		listenAddress: ":3000",
		handshaker:    p2p.NOPHandshakeFunc,
		deocder:       p2p.GoBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
