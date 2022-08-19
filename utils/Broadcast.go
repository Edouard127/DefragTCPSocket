package utils

import (
	"kamigen/socket/structs"
)

// BroadcastWorkers Broadcast message to all workers
func BroadcastWorkers(message []byte) {
	clients := structs.Clients
	for _, c := range clients {
		_, err := c.Conn.Write(message)
		if err != nil {
			return
		}
		c.Conn.Write([]byte{'\n'})
	}
}

// BroadcastListeners Broadcast message to all listeners
func BroadcastListeners(message []byte) {
	listeners := &structs.Listeners
	for _, c := range *listeners {
		_, err := c.Conn.Write(message)
		if err != nil {
			return
		}
		c.Conn.Write([]byte{'\n'})
	}
}
