package funcs

import (
	"encoding/binary"
	"fmt"
	"kamigen/socket/enums"
	. "kamigen/socket/structs"
	"kamigen/socket/utils"
	"net"
	"os"
)

func HandleServerCommand(connection *net.Conn, command *Command) {
	switch command.Header.GetCommand() {
	case EXIT:
		os.Exit(0)
	case ADD_WORKER:
		id := utils.RandomUint16(16)
		Clients = append(Clients, &Client{Player: binary.BigEndian.Uint16(id[:]), Conn: *connection})
		utils.LogFile(true, enums.INFO, fmt.Sprintf("Worker %d added", id))
		str := &AddWorkerResponse{
			Player{
				Player: id,
			},
		}
		// Send the worker the ID
		(*connection).Write(CreatePacket(ADD_WORKER, SERVER, str))
	/*case REMOVE_WORKER:
	for client := range Clients {
		if Clients[client].Player == binary.BigEndian.Uint16(command.Payload[:2]) {
			Clients = append(Clients[:client], Clients[client+1:]...)
		}
	}*/
	case ADD_LISTENER:
		Listeners = append(Listeners, &Listener{Conn: *connection})
		/*case REMOVE_LISTENER:
		for listener := range Listeners {
			if Listeners[listener].Conn == *connection {
				Listeners = append(Listeners[:listener], Listeners[listener+1:]...)
			}
		}*/
	}
}
