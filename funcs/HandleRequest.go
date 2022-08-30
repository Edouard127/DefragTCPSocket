package funcs

import (
	"github.com/fatih/color"
	"kamigen/socket/utils"
	"net"
	"strconv"
	"strings"
)

func HandleRequest(conn net.Conn, bSize int) {
	color.New(color.FgGreen, color.Bold).Println("New connection from:", conn.RemoteAddr().String())
	utils.LogFile(false, "New connection from:", conn.RemoteAddr().String())

	for {
		i, buffer, err := ReadAll(conn, bSize)
		if i == 0 {
			continue
		}
		if err != nil {
			utils.LogFile(false, "Connection from:", conn.RemoteAddr().String(), "closed")
			utils.LogFile(true, "Error: ", err.Error())
			color.New(color.FgRed, color.Bold).Println("Connection from:", conn.RemoteAddr().String(), "closed")
			conn.Close()
			return
		}
		c := strings.Fields(string(buffer))

		//length, _ := strconv.Atoi(c[0])
		if fragmented, err := strconv.Atoi(c[1]); err != nil {
			utils.LogFile(true, "Error: ", err.Error())
			conn.Close()
			return
		} else if fragmented == 1 || fragmented == 0 {
			go HandleCommand(&conn, &buffer, fragmented == 1)
		}
	}
}

// ReadAll Read all the data from the connection
func ReadAll(conn net.Conn, b int) (int, []byte, error) {
	buffer := make([]byte, b)
	i, err := conn.Read(buffer)
	if err != nil {
		return i, buffer, err
	}
	buffer = buffer[:i]
	return i, buffer, nil
}
