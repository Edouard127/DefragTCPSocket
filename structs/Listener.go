package structs

import "net"

type Listener struct {
	Conn net.Conn
}
