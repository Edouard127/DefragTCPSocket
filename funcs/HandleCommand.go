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

func HandleCommand(connection *net.Conn, command *[]byte) {
	con := *connection
	cmd := *command
	if len(cmd) == 0 {
		con.Write([]byte{structs.Packets["ERROR"]})
		con.Close()
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
	args := utils.GetArgs(request[1:])
	// TODO: Get more arguments from the request.
	intV, err := strconv.Atoi(request[0])
	// Store the command in a ClientCommands struct.
	cCom := structs.ClientCommand{Byte: byte(intV), Args: args}
	for _, v := range args {
		fmt.Println(string(v))
	}
	fmt.Println(intV, byte(intV), err)
	fmt.Println(cCom.GetPacketName(), cCom)
	message := utils.ByteArraysExtract(utils.GetArgs(request))

	switch cCom.Byte {
	case 0x05:
		// Register the client
		// Remove whitespaces from the name.
		name := strings.TrimSpace(string(args[0]))
		password := strings.TrimSpace(string(args[1]))
		client := structs.Client{Name: name, Conn: con, Password: password}
		structs.Clients = append(structs.Clients, &client)
		fmt.Println(client)
		_, err := con.Write(message)
		if err != nil {
			log.Fatal(err)
			return
		}
		con.Write([]byte{'\n'})
		break
	case 0x07:
		// Broadcast the message to all listeners.
		c := utils.GetClient(string(args[0]))
		_, e := c.Conn.Write(message)
		if e != nil {
			fmt.Println(e)
		}
		c.Conn.Write([]byte{'\n'})
	case 0x09:
		// Send chat message
		// Get the message from the arguments and add them to a byte array with a space between them.
		fmt.Println("Broadcasting:", string(message))
		c := utils.GetClient(string(args[0]))
		i, e := c.Conn.Write(message)
		if e != nil {
			fmt.Println(e)
		}
		c.Conn.Write([]byte{'\n'})

		fmt.Println("Bytes sent:", i)
		break
	case 0x0A:
		// Send chat message
		// Get the message from the arguments and add them to a byte array with a space between them.
		fmt.Println("Broadcasting:", string(message))
		c := utils.GetClient(string(args[0]))
		i, e := c.Conn.Write(message)
		if e != nil {
			fmt.Println(e)
		}
		c.Conn.Write([]byte{'\n'})

		fmt.Println("Bytes sent:", i)
		break
	}
}
