package utils

import (
	"kamigen/socket/structs"
)

// BroadcastWorkers Broadcast message to all workers
func BroadcastWorkers(message []byte) {
	clients := structs.Clients
	for client := range clients {
		_, err := clients[client].Conn.Write(message)
		if err != nil {
			return
		}
		clients[client].Conn.Write([]byte{'\n'})
	}
}

// BroadcastListeners Broadcast message to all listeners
func BroadcastListeners(message []byte) {
	listeners := structs.Listeners
	for client := range listeners {
		_, err := listeners[client].Conn.Write(message)
		if err != nil {
			return
		}
		listeners[client].Conn.Write([]byte{'\n'})
	}
}
