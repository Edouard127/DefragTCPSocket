
# Defrag TCP Socket


## Installation


```bash
  git clone https://github.com/Edouard127/DefragTCPSocket.git
  cd DefragTCPSocket
  npm run init
  npm run build
  npm run start
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


## Useful informations

#### Data transfers

Data sent through the socket is sent to each workers connected to the socket.

The data sent and received is encoded in Base64 format

#### Register a new worker

They commands are listed in the file `commands.ts`

Commands:

- 0 Exit and kill all the workers processes
- 1 Heartbeat. This will make sure that all the workers are connected and respond to the socket. If they don't respond to the socket within 20 seconds, the socket will be closed.
- 2 Login. This will send a server IP or hostname with the port to all the workers connected to the socket. (WIP)
- 3 Logout. This will make all the workers logout from the server. (WIP)
- 4 Add a new worker. This will register a new worker, if the password is not the same as the others, the connection will be closed with the code -1, they will receive a heartbeat every 5 seconds.
- 5 Remove a worker. This will remove the worker from the socket. (WIP)
- 6 Send a message in the connected server.
- 7 Baritone commands. The Baritone commands must match the commands registered in the file `isValidBaritone.ts`, [command name]: [command name, maximum amount of arguments allowed], if the number of arguments is greater than the maximum, the connection will be closed with the code 6.
- 8 Lambda commands. Will execute the command given. (WIP)

## Examples

Using the client.js

```bash
node build/src/client.js
[command] [password] [arguments] (6 password1234 Hello chat)
```

Video:
[![video](https://img.youtube.com/vi/jHUyMA1qg6U/maxresdefault.jpg)](https://www.youtube.com/watch?v=jHUyMA1qg6U)
