package p2p

const (
	IncomingMessage = 0x1
	IncomingStream  = 0x2
)

// RPC holds represents any arbitrary data that is being sent over
// each transport between two node in the network
type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
