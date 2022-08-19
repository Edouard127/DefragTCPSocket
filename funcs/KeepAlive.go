package funcs

import (
	"kamigen/socket/structs"
	"time"
)

// KeepAlive Keep clients alive
func KeepAlive() {
	for {
		clients := &structs.Clients
		for client := range *clients {
			(*clients)[client].Conn.SetDeadline(time.Now().Add(time.Second * 5))
			_, err := (*clients)[client].Conn.Write([]byte{'\n'})
			if err != nil {
				*clients = append((*clients)[:client], (*clients)[client+1:]...)
				return
			}
		}
		time.Sleep(time.Second * 5)
	}
}
