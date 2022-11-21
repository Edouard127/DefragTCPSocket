package structs

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Client struct {
	Player uint16
	Conn   net.Conn
}

type CommandHeader struct {
	Command     [1]byte
	Destination [1]byte
}

func (c *CommandHeader) GetCommand() Packet {
	return Packet(c.Command[0])
}

func (c *CommandHeader) GetDestination() Destination {
	return Destination(c.Destination[0])
}

type Command struct {
	Header  CommandHeader
	Payload []byte
}

func (c *Command) Read(buffer *bytes.Buffer) error {
	// Create the header buffer
	c.Header = CommandHeader{}
	if err := binary.Read(buffer, binary.LittleEndian, &c.Header); err != nil {
		return err
	}
	// Create the args buffer
	args := make([]byte, 0)
	if err := binary.Read(buffer, binary.LittleEndian, &args); err != nil {
		return err
	}
	c.Payload = args
	return nil
}
