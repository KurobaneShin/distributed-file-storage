package p2p

// RPC holds represents any arbitrary data that is being sent over
// each transport between two node in the network
type RPC struct {
	From    string
	Payload []byte
}
