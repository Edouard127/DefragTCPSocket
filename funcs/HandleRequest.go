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
		if i == 0 {
			continue
		}

		buffer = buffer[:i]

		println("Buffer size:", i)

		c := strings.Fields(string(buffer))

		//length, _ := strconv.Atoi(c[0])
		fragmented, err := strconv.Atoi(c[1])
		if err != nil {
			color.New(color.FgRed, color.Bold).Println("Error:", err)
			return
		}

		go HandleCommand(&conn, &buffer, fragmented == 1)
	}
}
