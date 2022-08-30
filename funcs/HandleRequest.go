package funcs

import (
	"kamigen/socket/enums"
	"kamigen/socket/utils"
	"net"
	"strconv"
	"strings"
)

func HandleRequest(conn net.Conn, bSize int) {
	utils.LogFile(false, enums.INFO, "New connection from:", conn.RemoteAddr().String())

	for {
		i, buffer, err := ReadAll(conn, bSize)
		if i == 0 {
			continue
		}
		if err != nil {
			utils.LogFile(false, enums.INFO, "Connection from:", conn.RemoteAddr().String(), "closed")
			utils.LogFile(true, enums.ERROR, "Error: ", err.Error())
			conn.Close()
			return
		}
		c := strings.Fields(string(buffer))

		//length, _ := strconv.Atoi(c[0])
		if fragmented, err := strconv.Atoi(c[1]); err != nil {
			utils.LogFile(true, enums.ERROR, "Error: ", err.Error())
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
	if i, err := conn.Read(buffer); err != nil {
		return i, buffer, err
	} else {
		buffer = buffer[:i]
		return i, buffer, nil
	}
}
