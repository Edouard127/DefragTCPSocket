package funcs

import (
	"github.com/fatih/color"
	"net"
	"strconv"
	"strings"
)

func HandleRequest(conn net.Conn) {
	color.New(color.FgGreen, color.Bold).Println("New connection from:", conn.RemoteAddr().String())
	for {
		buffer := make([]byte, 1024)

		i, _ := conn.Read(buffer)

		buffer = buffer[:i]

		println("Buffer size:", i)

		c := strings.Fields(string(buffer))

		//length, _ := strconv.Atoi(c[0])
		fragmented, _ := strconv.Atoi(c[1])

		go HandleCommand(&conn, &buffer, fragmented == 1)
	}
}
