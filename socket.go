package main

import (
    "fmt"
    "net"
    "os"
	"log"
	"github.com/fatih/color"
	"strings"
)


const (
    HOST = "localhost"
    PORT = "1984"
    TYPE = "tcp"
)

func main() {
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()

	color.Cyan("Listening on " + HOST + ":" + PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }

	color.Green("Connection etablished to: " + conn.RemoteAddr().String())

	// Store the request data in a splited array.
	request := strings.Split(string(buffer), " ")
	// Store the arguments of the request
	args := getArgs(request[1:])
	// Store the command in a ClientCommands struct.
	command := ClientCommands{toByte(request[0]), args}

	switch GetNumber(command.Byte) {
		case 0x04:
			client := Client{getString(args[0]), conn, getString(args[1])}
			clients = append(clients, &client)
			fmt.Println(client)
			conn.Write([]byte{Packets["OK"]})
		break
	}

}

func broadcast(message []byte) {
	for client := range clients {
		clients[client].Conn.Write(message)
	}
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
func GetNumber(s byte) int {
	return int(s - 0x30)
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

var Packets = map[string]byte {
	"EXIT": 					0x00, // user->server->client Notifies the client that the server is closing the connection.
	"OK": 						0x01, // client<->server Notifies the client that the server is ready to receive the next packet.
	"HEARTBEAT": 				0x02, // client<->server Ping packet.
	"LOGIN": 					0x03, // user->server<->client Notifies the server that the client is trying to login.
	"LOGOUT": 					0x04, // user->server<->client Notifies the server that the client is trying to logout.
	"ADD_WORKER": 				0x05, // client<->server Notifies the server of a new worker.
	"REMOVE_WORKER": 			0x06, // user<->server<->client Notifies the server that a worker has been removed.
	"GET_WORKERS": 				0x07, // user<->server<->client Notifies the server that the user wants to get the list of workers.
	"GET_WORKER_STATUS": 		0x08, // user<->server<->client Notifies the server that the user wants to get the status of a worker.
	"CHAT": 					0x09, // user->server<->client Notifies the server that the user wants to send a chat message.
	"BARITONE": 				0xA0, // user->server<->client Notifies the server that the user wants to send a baritone command.
	"LAMBDA": 					0xA1, // user->server<->client Notifies the server that the user wants to send a lambda command.
	"ERROR": 					0xA2, // client<->server<->user Notifies the user that the server or the client has encountered an error.
}


type ServerResponse struct {
	// Data of the packet.
	Data []byte
}
type ClientCommands struct {
	// The byte of the command.
	Byte byte
	// The arguments of the command.
	Args [][]byte
}

// Get the packet name from the ClientCommands byte.
func (c *ClientCommands) GetPacketName() string {
	for k, v := range Packets {
		if v == c.Byte {
			return k
		}
	}
	return "ERROR"
}
	



	