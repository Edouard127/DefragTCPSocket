package structs

var Listeners []*Listener

// BroadcastListeners Broadcast message to all listeners
func BroadcastListeners(message []byte) error {
	listeners := &Listeners
	for _, c := range *listeners {
		if _, err := c.Conn.Write(message); err != nil {
			return err
		} else if _, err := c.Conn.Write([]byte{'\n'}); err != nil {
			return err
		}
	}
	return nil
}
