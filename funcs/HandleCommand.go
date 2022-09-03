package funcs

import (
	"fmt"
	"kamigen/socket/enums"
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

	// Store the request data in a split array.
	request := strings.Fields(string(cmd))
	// Store the arguments of the request
	args := utils.GetArgs(request[4:])

	length, _ := strconv.Atoi(request[0])
	fragmentationByte, _ := strconv.Atoi(request[1])
	packetByte, _ := strconv.Atoi(request[2])
	packetFlag, _ := strconv.Atoi(request[3])

	// Store the command in a ClientCommands struct.
	c := structs.ClientCommand{Length: byte(length), Fragmented: byte(fragmentationByte), Byte: byte(packetByte), Flag: byte(packetFlag), Args: args}
	message := structs.ArgsExtract(utils.GetArgs(request))
	utils.LogFile(true, enums.INFO, "New command:", c.GetPacketName(), "args:", string(message))

	// TODO
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
					utils.LogFile(true, enums.INFO, "New listener registered with ID:", string(id))

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
					if err := structs.BroadcastWorkers(message); err != nil {
						utils.LogFile(false, enums.ERROR, "Error while broadcasting to workers:", err.Error())
						con.Write([]byte{structs.Packets["ERROR"]})
					}
				}
			case 0x06:
				{
					// Remove the client from the clients array.
					name := strings.TrimSpace(string(args[0]))
					clients := &structs.Clients
					if i, _, err := structs.GetClient(name); err == nil {
						*clients = append((*clients)[:i], (*clients)[i+1:]...)
						if err := structs.BroadcastWorkers(message); err != nil {
							utils.LogFile(false, enums.ERROR, "Error while broadcasting to workers:", err.Error())
							con.Write([]byte{structs.Packets["ERROR"]})
						}
					} else {
						utils.LogFile(false, enums.ERROR, "Client not found")
						con.Write([]byte{structs.Packets["ERROR"]})
					}
				}
			}
		}
	case 0x01:
		{
			utils.LogFile(false, enums.INFO, "Received command:", c.GetPacketName())
			fmt.Println("Client side")
			structs.BroadcastListeners(message)
		}
	case 0x02:
		{
			if err := structs.BroadcastWorkers(message); err != nil {
				utils.LogFile(false, enums.ERROR, "Error while broadcasting to workers:", err.Error())
				con.Write([]byte{structs.Packets["ERROR"]})
			}
		}
	case 0x03:
		{
			// Client & Game side
		}
	}
}
