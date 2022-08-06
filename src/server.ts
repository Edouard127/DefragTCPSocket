import net, { Socket } from "net"
import { bits } from "./utils/commands.js"
import { isValidBaritone } from "./utils/isValidBaritone"
import { responses, ServerResponses } from "./utils/serverResponses"
import zlib from "node:zlib"
import crypto from "crypto"
import "colors"
import Timeout from "./interfaces/Timeout.js"

const PORT = 1984
const HOST = "0.0.0.0"

const connectedSockets = new Map<Socket, String>()

const heartBeats = new Set<Timeout>()

const goodPass = (password: string) => {

    const keys = pKey(password)

    keys.forEach((k) => {
        if (k[1] === password) {
            return true
        }
    })
    return false
}

const pKey = (p: string) => [...connectedSockets.entries()].filter(v => v[1] === p).map((k) => k);

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
        const o = "PING "+ code
        write(o, socket)
    })
}

const killInactive = () => {
        heartBeats.forEach((v, k) => {
            if (new Date().getTime() - v.TIME > 10000 ) {
                end(responses.TIMEOUT, k.PONG.socket)
            }
        })
}

const write = (data: any, socket: Socket) => socket.write(zlib.deflateSync(data));

const end = (data: ServerResponses, socket: Socket) => socket.end(zlib.deflateSync(data+"\r\n"));


setInterval(() => {
    performKeepAlive()
    killInactive()
}, 10000)


const server = net.createServer((socket) => {
    socket.on("data", (data) => {
            console.log(data.toString())

            const command = zlib.inflateSync(data.buffer).toString()

            if (!command) return socket.end(responses.BAD_COMMAND)

            const parsed = [...JSON.parse(command)]

            const args = [...parsed].slice(2)
            console.log(parsed, args)
            
            switch(true) {
                case parsed[0] == bits.EXIT && goodPass(parsed[1]): connectedSockets.clear(); return kill(responses.DISCONNECT)
                case parsed[0] == bits.HEARTBEAT: console.log("Heartbeat received".bgGreen.white, parsed[1]); return heartBeats.delete(parsed[1]) ? 1 : socket.end(responses.DISCONNECT)
                case parsed[0] == bits.LOGIN: return broadcast([parsed[0], ...parsed.splice(2, 2)], socket, parsed[1])

                case parsed[0] == bits.ADD_WORKER: connectedSockets.set(socket, crypto.createHash('sha256').update(parsed[1]).digest('base64')); return write(responses.OK, socket)
                case parsed[0] == bits.REMOVE_WORKER: return write(connectedSockets.delete(socket) ? responses.OK : responses.WORKER_NOT_FOUND, socket)
                case parsed[0] == bits.CHAT: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                case parsed[0] == bits.BARITONE: return isValidBaritone(args) ? broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1]) : write(responses.BAD_ARGUMENTS, socket)
                case parsed[0] == bits.MOD_COMMAND: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                default: return end(responses.BAD_COMMAND, socket)
            }

    })
    socket.on("error", (e) => {
        console.log(e)
        kill(responses.SERVER_ERROR)
    })

}).listen(PORT, HOST)

server.on("connection", (socket) => {
    console.log("Connection established to ".bgGreen.white, socket.remoteAddress)
})

process.on("SIGINT", () =>  { kill(responses.SERVER_CLOSED); process.exit(0) })


