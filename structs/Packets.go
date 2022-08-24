package structs

var Packets = map[string]byte{
	"EXIT":            0x00, // user->server->client Notifies the client that the server is closing the connection.
	"OK":              0x01, // client<->server Notifies the client that the server is ready to receive the next packet.
	"HEARTBEAT":       0x02, // client<->server Ping packet.
	"LOGIN":           0x03, // user->server<->client Notifies the server that the client is trying to log in.
	"LOGOUT":          0x04, // user->server<->client Notifies the server that the client is trying to log out.
	"ADD_WORKER":      0x05, // user<->server Notifies the server of a new worker.
	"REMOVE_WORKER":   0x06, // user<->server Notifies the server that a worker has been removed.
	"GET_WORKERS":     0x07, // user<->server<->client Notifies the server that the user wants to get the list of workers.
	"JOB":             0x08, // user<->server<->client Notifies the server that the user wants to get the status of a worker.
	"CHAT":            0x09, // user->server<->client Notifies the server that the user wants to send a chat message.
	"BARITONE":        0x0A, // user->server<->client Notifies the server that the user wants to send a baritone command.
	"LAMBDA":          0x0B, // user->server<->client Notifies the server that the user wants to send a lambda command.
	"ERROR":           0x0C, // client<->server<->user Notifies the user that the server or the client has encountered an error.
	"LISTENER_ADD":    0x0D, // user<->server Notifies the server that a listener has been added.
	"LISTENER_REMOVE": 0x0E, // user<->server Notifies the server that a listener has been removed.
	"HIGHWAY_TOOLS":   0x0F, // user<->server<->client Notifies the server that the user wants to send a highwaytools command.
	"SCREENSHOT":      0x10, // user<->server<->client Notifies the server that the user wants to get a screenshot.
}
