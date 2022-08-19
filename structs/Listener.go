package structs

import "net"

type Listener struct {
	// The hash of the listener.
	Hash []byte
	// The type of the listener.
	// Type string
	// The status of the listener.
	// Status string
	// The connection of the listener.
	Conn net.Conn
}
