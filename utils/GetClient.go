package utils

import (
	"kamigen/socket/structs"
)

// GetClient Get pointer of struct by name
func GetClient(name string) structs.Client {
	clients := structs.Clients
	for _, v := range clients {
		if v.Name == name {
			return *v
		}
	}
	return *&structs.Client{}
}
