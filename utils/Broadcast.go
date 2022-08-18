package utils

import (
	"kamigen/socket/structs"
)

func Broadcast(message []byte) {
	clients := structs.Clients
	for client := range clients {
		_, err := clients[client].Conn.Write(message)
		if err != nil {
			return
		}
	}
}
