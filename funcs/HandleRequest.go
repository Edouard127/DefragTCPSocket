package funcs

import (
	"github.com/fatih/color"
	"kamigen/socket/structs"
	"log"
	"net"
)

func HandleRequest(conn net.Conn) {
	color.New(color.FgGreen, color.Bold).Println("New connection from:", conn.RemoteAddr().String())
	listeners := structs.Listeners
	listeners = append(listeners, &conn)
	for {
		buffer := make([]byte, 1024)

		i, err := conn.Read(buffer)
		buffer = buffer[:i]
		if err != nil {
			log.Fatal(err)
		}
		go HandleCommand(&conn, &buffer)
	}
}
