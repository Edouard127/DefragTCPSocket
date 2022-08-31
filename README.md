
# Lambda Remote Control Socket

A few months ago, I needed a simple solution to build highways on anarchy Minecraft servers.

Despite my short research, I didn't find anything, so...

I decided to make something useful for more than one person.


## Badges
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)\
[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.gg/J23U4YEaAr)
![Latest version](https://img.shields.io/github/tag/Edouard127/DefragTCPSocket?label=Latest)
![Downloads](https://img.shields.io/github/downloads/Edouard127/DefragTCPSocket/total)
![Maintainer](https://img.shields.io/badge/maintainer-Edouard127-blue)
![Activity](https://img.shields.io/github/commit-activity/w/Edouard127/DefragTCPSocket)
![Size](https://img.shields.io/github/languages/code-size/Edouard127/DefragTCPSocket)
![License](https://img.shields.io/github/license/Edouard127/DefragTCPSocket)

## [Plugin repository](https://github.com/Edouard127/LambdaRemoteControl)

## Features

- Multi-worker support
- WIP Get screenshot from the workers
- Login & Logout from servers
- Password secured connections
- Baritone & Lambda commands
- HighwayTools support

## FAQ

#### Can my workers be hijacked?

If you use a strong password, or the default randomly generated password, it is very unlikely that your workers will be intercepted by others.

## Installation

If you have already installed golang, you can skip this step. \
Install golang : https://go.dev/dl/ \
Add go to PATH : https://golang.org/doc/install/source#environment

```bash
git clone https://github.com/Edouard127/DefragTCPSocket.git
cd DefragTCPSocket
go build .
```
## Documentation

#### Requests

Any packet sent must have this:

TOTAL length of the packet\
Fragmentation byte: 0 or 1\
Packet byte\
Flag byte\
Args: Array of array of bytes

```go
type ClientCommand struct {
	// The length of the command.
	Length byte
	// Fragmentation byte.
	Fragmented byte
	// The byte of the command.
	Byte byte
	// The flags of the command.
	Flag byte
	// The arguments of the command.
	Args [][]byte
}
```

#### Destination Flags
```go
var Flags = map[string]byte{
	"SERVER": 0x00, // Server
	"CLIENT": 0x01, // Listeners
	"GAME":   0x02, // Workers
	"BOTH":   0x03, // CLIENT & GAME
}
```

#### Commands
The commands are arrays of bytes
They are at index 4 of the packet

#### Workers

A worker can be registered via the packet `5`\
`ADD_WORKER {Length Fragmented Byte Flag [Username Password]}`\
Example: \
`ADD_WORKER {44 0 5 0 [[75 97 109 105 103 101 110] [49 53 102 57 57 54 48 53 45 51 52 102 101 45 52 100 55 53 45 56 54 56 102 45 57 55 100 54 98 97 99 50 101 102 50 102]]}`

A worker can also be removed with the packet `6`\
`REMOVE_WORKER {Length Fragmented Byte Flag [Username Password]}`\
Example: \
`REMOVE_WORKER {44 0 6 0 [[75 97 109 105 103 101 110] [49 53 102 57 57 54 48 53 45 51 52 102 101 45 52 100 55 53 45 56 54 56 102 45 57 55 100 54 98 97 99 50 101 102 50 102]]}`


#### Job Tracking
When sending a: Baritone or HighwayTools packet, a job will be created.\
A job will emit his status when the game receives a:
- JobEvent
- StartPathingEvent
- StopPathingEvent
- UpdatePathingEvent

Is emitted

```
Started:
83 0 8 1 Job type:1 Status:0 Goal:BetterBlockPos{x=0,y=0,z=0} Player:Kamigen Scheduled


Finished:
83 0 8 1 Job type:1 Status:1 Goal:BetterBlockPos{x=0,y=0,z=0} Player:Kamigen Scheduled
```
*Take note that the format can change based on conditions*

Worker type:
```kotlin
enum class EWorkerType(val byte: Int) {
    HIGHWAY(byte = 0x00),
    BARITONE(byte = 0x01),
}
```
Worker status
```kotlin
enum class EWorkerStatus(val byte: Byte) {
    BUSY(byte = 0x00),
    IDLE(byte = 0x01),
}
```

If the worker is stuck, the game will send emit an event with the `JOB_STUCK` byte.\
It will also be sent to the socket with the flag `1`

Job events:
```kotlin
enum class EJobEvents(val byte: Int) {
    JOB_STARTED(byte = 0x00),
    JOB_FINISHED(byte = 0x01),
    JOB_FAILED(byte = 0x02),
    JOB_PAUSED(byte = 0x03),
    JOB_RESUMED(byte = 0x04),
    JOB_CANCELLED(byte = 0x05),
    JOB_SCHEDULED(byte = 0x06),
    JOB_STUCK(byte = 0x07),
}
```

#### TODO: Remaining

```go
var Packets = map[string]byte{
	"EXIT":            0x00, // user->server->client Notifies the client that the server is closing the connection.
	"OK":              0x01, // client<->server Notifies the client that the server is ready to receive the next packet.
	"HEARTBEAT":       0x02, // client<->server Ping packet.
	"LOGIN":           0x03, // user->server<->client Notifies the server that the client is trying to log in.
	"LOGOUT":          0x04, // user->server<->client Notifies the server that the client is trying to log out.
	"ADD_WORKER":      0x05, // user<->server Notifies the server of a new worker.
	"REMOVE_WORKER":   0x06, // user<->server Notifies the server that a worker has been removed.
	"INFORMATIONS":    0x07, // user<->server<->client Notifies the server that the user wants to get the information of a worker.
	"JOB":             0x08, // user<->server<->client Notifies the server that the user wants to get the status of a worker.
	"CHAT":            0x09, // user->server<->client Notifies the server that the user wants to send a chat message.
	"BARITONE":        0x0A, // user->server<->client Notifies the server that the user wants to send a baritone command.
	"LAMBDA":          0x0B, // user->server<->client Notifies the server that the user wants to send a lambda command.
	"ERROR":           0x0C, // client<->server<->user Notifies the user that the server or the client has encountered an error.
	"LISTENER_ADD":    0x0D, // user<->server Notifies the server that a listener has been added.
	"LISTENER_REMOVE": 0x0E, // user<->server Notifies the server that a listener has been removed.
	"HIGHWAY_TOOLS":   0x0F, // user<->server<->client Notifies the server that the user wants to send a highwaytools command.
	"SCREENSHOT":      0x10, // user<->server<->client Notifies the server that the user wants to get a screenshot.
	"GET_JOBS":        0x11, // user<->server<->client Notifies the server that the user wants to get the list of jobs.
	"ROTATE":          0x12, // user<->server<->client Rotates the worker head position.
}
```

Latest dev footage: https://youtu.be/j80Uqv2IxQI

Packages graph:

![graph](./godepgraph.png)

