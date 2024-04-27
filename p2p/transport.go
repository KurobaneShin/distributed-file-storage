package p2p

// Peer is an interface that represents the remote node
type Peer interface{}

// transport is anything that handles communication beetween the nodes in the network
// this can be of the form (TCP, UDP, websockets, etc)
type Transport interface{}
