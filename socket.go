package main

import (
	"fmt"
	"github.com/fatih/color"
	"kamigen/socket/funcs"
	"net"
)

const (
	HOST        = "localhost"
	PORT        = "1984"
	BUFFER_SIZE = 1024
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
		go funcs.KeepAlive()
		// Execute the handleRequest function in a new goroutine
		go funcs.HandleRequest(conn, BUFFER_SIZE)
	}
}
