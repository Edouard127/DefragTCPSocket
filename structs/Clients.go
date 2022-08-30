package structs

import "errors"

var Clients []*Client

// GetClient Get pointer of struct by name
func GetClient(name string) (*Client, error) {
	clients := &Clients
	for _, v := range *clients {
		if v.Name == name {
			return v, nil
		}
	}
	return &Client{}, errors.New("client not found")
}

// BroadcastWorkers Broadcast message to all workers
func BroadcastWorkers(message []byte) error {
	clients := &Clients
	for _, c := range *clients {
		if _, err := c.Conn.Write(message); err != nil {
			return err
		} else if _, err := c.Conn.Write([]byte{'\n'}); err != nil {
			return err
		}
	}
	return nil
}
