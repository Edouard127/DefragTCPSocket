package structs

var Flags = map[string]byte{
	"SERVER": 0x00, // Server
	"CLIENT": 0x01, // Listeners
	"GAME":   0x02, // Workers
	"BOTH":   0x03, // CLIENT & GAME
}
