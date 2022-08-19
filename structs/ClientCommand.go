package structs

type ClientCommand struct {
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
