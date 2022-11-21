package main

import (
	"fmt"
	"kamigen/socket/enums"
	"kamigen/socket/funcs"
	"kamigen/socket/utils"
	"net"
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
			panic(err)
		}
	}(listener)
	fmt.Printf("Listening on %s:%s\n", HOST, PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.LogFile(true, enums.ERROR, "Error accepting: ", err.Error())
		}

		// Execute the keepAlive function in a new goroutine
		//go funcs.KeepAlive()
		// Execute the handleRequest function in a new goroutine
		go funcs.HandleRequest(conn)
	}
}
