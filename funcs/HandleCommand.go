package funcs

import (
	"fmt"
	"kamigen/socket/structs"
	"kamigen/socket/utils"
	"log"
	"net"
	"strconv"
	"strings"
)

func HandleCommand(connection *net.Conn, command *[]byte, needFragmentation bool) {
	con := *connection
	cmd := *command

	if len(cmd) == 0 {
		_, write := con.Write([]byte{structs.Packets["ERROR"]})
		if write != nil {
			return
		}
		err := con.Close()
		if err != nil {
			return
		}
	}

	/*
		TODO Check if the command is valid.
		TODO Check if the arguments are valid.
		TODO: Check if the command is valid.
			if int(getByteNumber(toByte(request[1]))) != len(command.Args) {
				conn.Write([]byte{Packets["ERROR"]})
				conn.Close()
				return
			}
			if command.GetPacketName() == "ERROR" {
				conn.Write([]byte{Packets["ERROR"]})
				conn.Close()
				return
			}
	*/

	// Store the request data in a splited array.
	request := strings.Fields(string(cmd))
	// Store the arguments of the request
	args := utils.GetArgs(request[4:])

	length, _ := strconv.Atoi(request[0])
	fragmentationByte, _ := strconv.Atoi(request[1])
	packetByte, _ := strconv.Atoi(request[2])
	packetFlag, _ := strconv.Atoi(request[3])

	// Store the command in a ClientCommands struct.
	c := structs.ClientCommand{Length: byte(length), Fragmented: byte(fragmentationByte), Byte: byte(packetByte), Flag: byte(packetFlag), Args: args}
	message := utils.ByteArraysExtract(utils.GetArgs(request))

	// If it needs fragmentation, loop and add the data to the buffer
	if needFragmentation {
		buffer := make([]byte, c.Length)
		for {
			if c.Byte == 0 {
				break
			}
			buffer = append(buffer, cmd...)
		}
	}

	switch c.Flag {
	case 0x00:
		{
			// Server side
			switch c.Byte {
			case 0x0D:
				{
					// Register the listener
					listeners := &structs.Listeners
					// Get random bytes for the listener id.
					id := utils.GetRandomBytes(16)

					fmt.Println("Registering listener with id:", id, "(", string(id), ")")

					listener := structs.Listener{Hash: id, Conn: con}

					*listeners = append(*listeners, &listener)
				}
			case 0x05:
				{
					// Register the client
					// Remove whitespaces from the name.
					name := strings.TrimSpace(string(args[0]))
					password := strings.TrimSpace(string(args[1]))
					client := structs.Client{Name: name, Conn: con, Password: password}
					clients := &structs.Clients
					*clients = append(*clients, &client)
					_, err := con.Write(message)
					if err != nil {
						log.Fatal(err)
						return
					}
					_, err = con.Write([]byte{'\n'})
					if err != nil {
						return
					}
				}
			}
		}
	case 0x01:
		{
			utils.LogFile(false, "[INFO]", "Received command:", c.GetPacketName())
			fmt.Println("Client side")
			structs.BroadcastListeners(message)
		}
	case 0x02:
		{
			// Game side
			c, err := structs.GetClient(string(args[0]))
			if err != nil {
				con.Write([]byte{structs.Packets["ERROR"]})
				fmt.Println(err)
				return
			}
			i, e := c.Conn.Write(message)
			if e != nil {
				fmt.Println(e)
			}
			_, err = c.Conn.Write([]byte{'\n'})
			if err != nil {
				return
			}

			fmt.Println("Bytes sent:", i)
		}
	case 0x03:
		{
			// Client & Game side
		}
	}
}
