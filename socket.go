package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net"
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

	// Store the request data in a splited array.
	request := strings.Fields(string(buffer))
	// Store the arguments of the request
	args := getArgs(request[4:])
	// TODO: Get more arguments from the request.
	// Store the command in a ClientCommands struct.
	command := ClientCommand{toByte(request[0]), args}
	fmt.Println(getByteNumber(toByte(request[1])), getByteNumber(toByte(request[2])), getByteNumber(toByte(request[3])))
	fmt.Println(command.GetPacketName(), command)
	fmt.Println(int(getByteNumber(toByte(request[1]))), len(command.Args))

	// Check if the packet is valid.
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

	switch getByteNumber(command.Byte) {
	case 0x05:
		// Register the client
		client := Client{getString(args[0]), conn, getString(args[1])}
		clients = append(clients, &client)
	case 0x06:
		// Remove client
		for i, v := range clients {
			if v.Name == getString(args[0]) {
				clients = append(clients[:i], clients[i+1:]...)
			}
		}
	default:
		for i, v := range clients {
			if v.Name == getString(args[0]) {
				clients[i].Conn.Write(encode(command))
			}
		}

	}

}

// Keep clients alive
func keepAlive() {
	for {
		for client := range clients {
			clients[client].Conn.SetDeadline(time.Now().Add(time.Second * 5))
			_, err := clients[client].Conn.Write([]byte{Packets["KEEPALIVE"]})
			if err != nil {
				clients = append(clients[:client], clients[client+1:]...)
				return
			}
		}
		time.Sleep(time.Second * 10)
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

func encode(data interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func toByte(s string) byte {
	return byte(s[0])
}

func toByteArray(s []string) []byte {
	var b []byte
	for _, v := range s {
		b = append(b, toByte(v))
	}
	return b
}

func getArgs(args []string) [][]byte {
	var b [][]byte
	for _, v := range args {
		b = append(b, []byte(v))
	}
	return b
}
func getString(args []byte) string {
	var s string
	for _, v := range args {

		s += string(v)
	}
	return s
}

// Get number from byte
func getByteNumber(s byte) byte {
	return s - 0x30
}

var clients []*Client

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

type ServerResponse struct {
	// Data of the packet.
	Data []byte
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
		if v == getByteNumber(c.Byte) {
			return k
		}
	}
	return "ERROR"
}
