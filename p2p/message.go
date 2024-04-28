package p2p

import "net"

// Message holds represents any arbitrary data that is being sent over
// each transport between two node in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
