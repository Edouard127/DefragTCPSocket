package funcs

import (
	"fmt"
	"kamigen/socket/enums"
	. "kamigen/socket/structs"
	"kamigen/socket/utils"
	"net"
)

func HandleCommand(connection *net.Conn, command *Command) {
	utils.LogFile(true, enums.INFO, fmt.Sprintf("Command received from %s: %s", (*connection).RemoteAddr().String(), command.Header.GetCommand().String()))
	switch command.Header.GetDestination() {
	case SERVER:
		HandleServerCommand(connection, command)
	case WORKER:
		HandleWorkerCommand(command)
	}
}
