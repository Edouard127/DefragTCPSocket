import net, { Socket } from "net"
import { bits } from "./utils/commands.js"
import { isValidBaritone } from "./utils/isValidBaritone"
import { responses, ServerResponses } from "./utils/serverResponses"
import crypto from "crypto"
import "colors"
import Timeout from "./interfaces/Timeout.js"

const PORT = 1984
const HOST = "0.0.0.0"

const connectedSockets = new Map<Socket, String>()

const heartBeats = new Set<Timeout>()

const goodPass = (password: string) => {
    console.log("Checking password".bgYellow.white)
    const keys = pKey(password)
    var valid = false
    keys.forEach((k) => {
        if (k[1] === crypto.createHash('sha256').update(password).digest('base64')) {
            valid = true
        }
    })
    return valid
}

const pKey = (p: string) => [...connectedSockets.entries()].filter(v => v[1] === crypto.createHash('sha256').update(p).digest('base64')).map((k) => k);

const broadcast = (data: Array<string>, sender: Socket, password: string) => {
    connectedSockets.forEach((_, socket) => {
        if (socket !== sender && goodPass(password)) {
            socket.setEncoding('utf8');
            write(data.toString().replace(/,/g, " "), socket);
        }
    })
}

const kill = (code: ServerResponses) => {
    console.log("Killing connections".bgRed.white)
    connectedSockets.forEach((_, socket) => end(code, socket))
}

const performKeepAlive = () => {
    connectedSockets.forEach((_, socket) => {
        const code = (Math.random() + 1).toString(16).substring(10);
        heartBeats.add({ PONG: { socket: socket, code: code }, TIME: new Date().getTime() })
        write("PING " + code, socket)
    })
}

const killInactive = () => {
    heartBeats.forEach((k) => {
        if (new Date().getTime() - k.TIME > 20000 ) {
            console.log("Killing inactive connection".bgRed.white)
            end(responses.TIMEOUT, k.PONG.socket)
        }
        heartBeats.delete(k)
    })
}

const write = (data: String, socket: Socket) => socket.write(Buffer.from(data).toString("base64")+"\r\n");

const end = (data: ServerResponses, socket: Socket) => socket.end(Buffer.from(data).toString("base64")+"\r\n");



setInterval(() => {
    performKeepAlive()
    killInactive()
}, 5000)


const server = net.createServer((socket) => {
    socket.on("data", (data) => {
            const command = Buffer.from(data.toString().trim(), "base64").toString("ascii")

            if (!command) return socket.end(responses.BAD_COMMAND)
            if (!Array.isArray(JSON.parse(command))) return end(responses.BAD_COMMAND, socket)

            const parsed = [...JSON.parse(command)]

            const args = [...parsed].slice(2)
            
            switch(true) {   
                case parsed[0] == bits.EXIT && goodPass(parsed[1]): connectedSockets.clear(); return kill(responses.DISCONNECT)
                case parsed[0] == bits.HEARTBEAT: console.log("Heartbeat received".bgGreen.white, parsed.join(" ")); return heartBeats.delete(parsed[0])
                case parsed[0] == bits.LOGIN: return broadcast([parsed[0], ...parsed.splice(2, 2)], socket, parsed[1])

                case parsed[0] == bits.ADD_WORKER: console.log("Worker added".bgGreen.white, parsed.join(" ")); connectedSockets.set(socket, crypto.createHash('sha256').update(parsed[1]).digest('base64')); return write(responses.OK, socket)
                case parsed[0] == bits.REMOVE_WORKER: return write(connectedSockets.delete(socket) ? responses.OK : responses.WORKER_NOT_FOUND, socket)
                case parsed[0] == bits.CHAT: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                case parsed[0] == bits.BARITONE: return isValidBaritone(args) ? broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1]) : write(responses.BAD_ARGUMENTS, socket)
                case parsed[0] == bits.MOD_COMMAND: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                default: return end(responses.BAD_COMMAND, socket)
            }
    })
    socket.on("close", () => {
        connectedSockets.delete(socket)
        heartBeats.clear()
        console.log("Connection closed".bgRed.white)
    })
    socket.on("error", () => {
        connectedSockets.delete(socket)
        heartBeats.clear()
        console.log("Connection closed".bgRed.white)
    })
    socket.on("error", () => kill(responses.SERVER_ERROR))

}).listen(PORT, HOST)

server.on("connection", (socket) => {
    console.log("Connection established to ".bgGreen.white, socket.remoteAddress)
})

process.on("SIGINT", () =>  { kill(responses.SERVER_CLOSED); process.exit(0) })


