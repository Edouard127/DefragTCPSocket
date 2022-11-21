package funcs

import (
	"bytes"
	"kamigen/socket/enums"
	. "kamigen/socket/structs"
	"kamigen/socket/utils"
	"net"
)

func HandleRequest(conn net.Conn) {
	utils.LogFile(true, enums.INFO, "New connection from:", conn.RemoteAddr().String())

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		utils.LogFile(true, enums.ERROR, "Error reading:", err.Error())
	}

	command := Command{}

	rBuffer := bytes.NewBuffer(buf[:reqLen])
	// Read the header
	err = command.Read(rBuffer)
	if err != nil {
		utils.LogFile(true, enums.ERROR, "Error reading:", err.Error())
	}
	command.Payload = rBuffer.Bytes()
	HandleCommand(&conn, &command)
}
