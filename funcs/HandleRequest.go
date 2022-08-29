package funcs

import (
	"github.com/fatih/color"
	"net"
)

func HandleRequest(conn net.Conn) {
	color.New(color.FgGreen, color.Bold).Println("New connection from:", conn.RemoteAddr().String())
	for {
		buffer := make([]byte, 1024)
		i, err := conn.Read(buffer)
		if err != nil {
			color.New(color.FgRed, color.Bold).Println("Connection from:", conn.RemoteAddr().String(), "closed")
			conn.Close()
			return
		}
		buffer = buffer[:i]
		go HandleCommand(&conn, &buffer, false)
	}
}
