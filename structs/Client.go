package structs

import (
	"net"
)

type Client struct {
	// The name of the client.
	Name string
	// The connection to the client.
	Conn net.Conn
	// The password of the client.
	Password string
}
