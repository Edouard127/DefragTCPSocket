package structs

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
	return string(rune(c.Length)) + " " + string(rune(c.Fragmented)) + " " + string(rune(c.Byte)) + " " + string(rune(c.Flag)) + " " + string(ArgsExtract(c.Args))
}
func ArgsExtract(arr [][]byte) []byte {
	var b []byte
	for _, v := range arr {
		b = append(b, v...)
		b = append(b, ' ')
	}
	return b
}
