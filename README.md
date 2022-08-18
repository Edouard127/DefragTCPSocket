
# Defrag TCP Socket Golang


## Installation

If you have already installed golang, you can skip this step. \
Install golang : https://go.dev/dl/ \
Add go to PATH : https://golang.org/doc/install/source#environment

```bash
git clone https://github.com/Edouard127/DefragTCPSocket.git
cd DefragTCPSocket
go build
```


## Features

- WIP Multi-worker support
- WIP Get screenshot from the workers
- WIP Login & Logout from servers
- Password secured connections
- Baritone & Lambda commands
- Send messages to servers


## FAQ

#### Can my workers be hijacked?

If you use a strong password, or the default randomly generated password, it is very unlikely that your workers will be intercepted by others.

#### Why my game is crashing ?

Probably because kotlin & java are not my preferred language & not the one I know the most

#### Why my server is crashing ?

Because I started using golang 2 days ago :trollface:


## Useful information

#### Data transfers

Data sent through the socket is sent to each worker connected to the socket.
The data sent is a encoded struct of ClientCommand:

```go
type ClientCommand struct {
	Command string
	Data    [][]byte
}
```


#### Register a new worker

```go
type Client struct {
	// The name of the client.
	Name    string
	// The connection to the client.
	Conn    net.Conn
	// The password of the client.
	Password string
}
```



Packets:

- 0x00 Exit and kill all the workers processes
- 0x01 OK Packet.
- 0x02 Heartbeat. This will make sure that all the workers are connected and respond to the socket. If they don't respond, the client will be removed
- 0x03 Login. This will send a server IP or hostname with the port to the specified workers connected to the socket. (WIP)
- 0x04 Logout. This will make all the workers logout from the server. (WIP)
- 0x05 Add a new worker. This will register a new worker, they will receive a heartbeat every 10 seconds.
- 0x06 Remove a worker. This will remove the worker from the socket. If the client is not found or the password does not match, the server will respond with 0x0C
- 0x07 Get a worker data. If the client is not found or the password does not match, the server will respond with 0x0C
- 0x08 Get workers status
- 0x09 Send a chat message
- 0x0A Send a baritone command
- 0x0B Send a lambda command
- 0x0C Error packet


#### Protocol

Each packet sent must match the hardcoded protocol, a documentation will soon be available

```
[Packet]
[Data...] The number of arguments
```




## Examples

Using the client.js

```bash
node client.js
[Packet] 5
[Data...] The arguments are strings splited into array of bytes
```

Video: SOON
