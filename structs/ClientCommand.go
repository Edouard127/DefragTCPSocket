package structs

import "kamigen/socket/utils"

type ClientCommand struct {
	// The length of the command.
	Length byte
	// Fragmentation byte.
	Fragmented byte
	// The byte of the command.
	Byte byte
	// The flags of the command.
	Flag byte
	// The arguments of the command.
	Args [][]byte
}

// GetPacketName Get the packet name from the ClientCommands byte.
func (c *ClientCommand) GetPacketName() string {
	for k, v := range Packets {
		if v == c.Byte {
			return k
		}
	}
	return "ERROR"
}

// GetCommand Get the command from the ClientCommands bytes.
func (c *ClientCommand) GetCommand() string {
	return string(c.Length) + " " + string(c.Fragmented) + " " + string(c.Byte) + " " + string(c.Flag) + " " + string(utils.ByteArraysExtract(c.Args))
}
