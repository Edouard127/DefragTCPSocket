export const bits = {
    EXIT:              "0", // user->server->client Notifies the client that the server is closing the connection.
	OK:                "1", // client<->server Notifies the client that the server is ready to receive the next packet.
	HEARTBEAT:         "2", // client<->server Ping packet.
	LOGIN:             "3", // user->server<->client Notifies the server that the client is trying to log in.
	LOGOUT:            "4", // user->server<->client Notifies the server that the client is trying to log out.
	ADD_WORKER:        "5", // user<->server Notifies the server of a new worker.
	REMOVE_WORKER:     "6", // user<->server Notifies the server that a worker has been removed.
	GET_WORKERS:       "7", // user<->server<->client Notifies the server that the user wants to get the list of workers.
	GET_WORKER_STATUS: "8", // user<->server<->client Notifies the server that the user wants to get the status of a worker.
	CHAT:              "9", // user->server<->client Notifies the server that the user wants to send a chat message.
	BARITONE:          "10", // user->server<->client Notifies the server that the user wants to send a baritone command.
	LAMBDA:            "11", // user->server<->client Notifies the server that the user wants to send a lambda command.
	ERROR:             "12", // client<->server<->user Notifies the user that the server or the client has encountered an error.
} as const;