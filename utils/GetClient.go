package utils

import (
	"kamigen/socket/structs"
)

// GetClient Get pointer of struct by name
func GetClient(name string) *structs.Client {
	for _, v := range structs.Clients {
		if v.Name == name {
			return v
		}
	}
	return &structs.Client{}
}
