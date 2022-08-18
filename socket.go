package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	HOST = "localhost"
	PORT = "1984"
)

func main() {
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)
	color.New(color.BgCyan, color.Bold).Println("Server started on:", HOST+":"+PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err, "error")
			continue
		}

		// Execute the keepAlive function in a new goroutine
		go keepAlive()
		// Execute the handleRequest function in a new goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	color.New(color.FgGreen, color.Bold).Println("New connection from:", conn.RemoteAddr().String())
	listeners = append(listeners, &conn)
	buffer := make([]byte, 1024)

	i, err := conn.Read(buffer)
	buffer = buffer[:i]
	if err != nil {
		log.Fatal(err)
	}

	if len(buffer) == 0 {
		conn.Write([]byte{Packets["ERROR"]})
		conn.Close()
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
	request := strings.Fields(string(buffer))
	// Store the arguments of the request
	args := getArgs(request[1:])
	// TODO: Get more arguments from the request.
	intV, err := strconv.Atoi(request[0])
	// Store the command in a ClientCommands struct.
	command := ClientCommand{byte(intV), args}
	for _, v := range args {
		fmt.Println(string(v))
	}
	fmt.Println(intV, byte(intV), err)
	fmt.Println(command.GetPacketName(), command)
	message := AArrayByteToArrByte(getArgs(request))

	switch command.Byte {
	case 0x05:
		// Register the client
		// Remove whitespaces from the name.
		name := strings.TrimSpace(string(args[0]))
		password := strings.TrimSpace(string(args[1]))
		client := Client{name, conn, password}
		clients = append(clients, &client)
		fmt.Println(client)
		_, err := conn.Write(message)
		if err != nil {
			log.Fatal(err)
			return
		}
		conn.Write([]byte{'\n'})
		break
	case 0x07:
		// Broadcast the message to all listeners.
		c := getClient(string(args[0]))
		_, e := c.Conn.Write(message)
		if e != nil {
			fmt.Println(e)
		}
		c.Conn.Write([]byte{'\n'})
	case 0x09:
		// Send chat message
		// Get the message from the arguments and add them to a byte array with a space between them.
		fmt.Println("Broadcasting:", string(message))
		c := getClient(string(args[0]))
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
		c := getClient(string(args[0]))
		i, e := c.Conn.Write(message)
		if e != nil {
			fmt.Println(e)
		}
		c.Conn.Write([]byte{'\n'})

		fmt.Println("Bytes sent:", i)
		break
	}

}

// Keep clients alive
func keepAlive() {
	for {
		for client := range clients {
			clients[client].Conn.SetDeadline(time.Now().Add(time.Second * 5))
			_, err := clients[client].Conn.Write([]byte{'\n'})
			if err != nil {
				clients = append(clients[:client], clients[client+1:]...)
				return
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func broadcast(message []byte) {
	for client := range clients {
		_, err := clients[client].Conn.Write(message)
		if err != nil {
			return
		}
	}
}

func broadcastListeners(message []byte) {
	for _, v := range listeners {
		_, err := (*v).Write(message)
		if err != nil {
			return
		}
		(*v).Write([]byte{'\n'})
	}
}

func getArgs(args []string) [][]byte {
	var b [][]byte
	for _, v := range args {
		b = append(b, []byte(v))
	}
	return b
}
func AArrayByteToArrByte(arr [][]byte) []byte {
	var b []byte
	for _, v := range arr {
		b = append(b, v...)
		b = append(b, ' ')
	}
	return b
}

var clients []*Client
var listeners []*net.Conn

// Get pointer of struct by name
func getClient(name string) *Client {
	for _, v := range clients {
		if v.Name == name {
			return v
		}
	}
	return &Client{}
}

type Client struct {
	// The name of the client.
	Name string
	// The connection to the client.
	Conn net.Conn
	// The password of the client.
	Password string
}

var Packets = map[string]byte{
	"EXIT":              0x00, // user->server->client Notifies the client that the server is closing the connection.
	"OK":                0x01, // client<->server Notifies the client that the server is ready to receive the next packet.
	"HEARTBEAT":         0x02, // client<->server Ping packet.
	"LOGIN":             0x03, // user->server<->client Notifies the server that the client is trying to log in.
	"LOGOUT":            0x04, // user->server<->client Notifies the server that the client is trying to log out.
	"ADD_WORKER":        0x05, // user<->server Notifies the server of a new worker.
	"REMOVE_WORKER":     0x06, // user<->server Notifies the server that a worker has been removed.
	"GET_WORKERS":       0x07, // user<->server<->client Notifies the server that the user wants to get the list of workers.
	"GET_WORKER_STATUS": 0x08, // user<->server<->client Notifies the server that the user wants to get the status of a worker.
	"CHAT":              0x09, // user->server<->client Notifies the server that the user wants to send a chat message.
	"BARITONE":          0x0A, // user->server<->client Notifies the server that the user wants to send a baritone command.
	"LAMBDA":            0x0B, // user->server<->client Notifies the server that the user wants to send a lambda command.
	"ERROR":             0x0C, // client<->server<->user Notifies the user that the server or the client has encountered an error.
}

type ClientCommand struct {
	// The byte of the command.
	Byte byte
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
