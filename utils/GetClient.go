package utils

import (
	"errors"
	"kamigen/socket/structs"
)

// GetClient Get pointer of struct by name
func GetClient(name string) (*structs.Client, error) {
	clients := &structs.Clients
	for _, v := range *clients {
		if v.Name == name {
			return v, nil
		}
	}
	return &structs.Client{}, errors.New("client not found")
}
