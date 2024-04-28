package p2p

import "net"

// RPC holds represents any arbitrary data that is being sent over
// each transport between two node in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
