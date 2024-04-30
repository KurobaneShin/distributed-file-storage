package p2p

// Peer is an interface that represents the remote node
type Peer interface {
	Close() error
}

// transport is anything that handles communication beetween the nodes in the network
// this can be of the form (TCP, UDP, websockets, etc)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
