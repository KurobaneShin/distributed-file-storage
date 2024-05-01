package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface {
	RemoteAddr() net.Addr
	Close() error
}

// transport is anything that handles communication beetween the nodes in the network
// this can be of the form (TCP, UDP, websockets, etc)
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
