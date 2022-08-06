import net, { Socket } from "net"
import { bits } from "./utils/commands.js"
import { isValidBaritone } from "./utils/isValidBaritone"
import { responses } from "./utils/serverResponses"
import zlib from "node:zlib"
import crypto from "crypto"
import "colors"

const PORT = 1984
const HOST = "0.0.0.0"

const connectedSockets = new Set<Socket>();

let pass = new Map<Socket, string>()

const goodPass = (socket: Socket, password: string) => {
    if (pass.has(socket)) {
        return pass.get(socket) === password
    }
    return false
}

const broadcast = (data: Array<string>, sender: Socket, password: string) => {
    connectedSockets.forEach((sock) => {
        sock = sock as Socket
        if (sock !== sender && goodPass(sock, password)) {
            sock.setEncoding('utf8');
            sock.write(data.toString().replace(/,/g, " ")+"\r\n");
        }
    })
}

const kill = (code: string) => {
    console.log("Killing connections".bgRed.white)
    connectedSockets.forEach((sock) => sock.end(code))
}

const server = net.createServer((socket) => {
    socket.on("data", (data) => {
            console.log(data.toString())

            const command = /*zlib.inflateSync(data.buffer).toString()*/data.toString()

            if (!command) return socket.end(responses.BAD_COMMAND)

            const parsed = [...JSON.parse(command)]

            const args = [...parsed].slice(2)
            console.log(parsed, args)
            
            switch(true) {
                case parsed[0] == bits.EXIT && goodPass(socket, parsed[1]): pass.clear(); return kill(responses.DISCONNECT)
                case parsed[0] == bits.LOGIN: return broadcast([parsed[0], ...parsed.splice(2, 2)], socket, parsed[1])

                case parsed[0] == bits.ADD_WORKER: pass.set(socket, crypto.createHash('sha256').update(parsed[1]).digest('base64')); return socket.write(responses.OK)
                case parsed[0] == bits.REMOVE_WORKER: return socket.write(pass.delete(socket) ? responses.OK : responses.WORKER_NOT_FOUND)
                case parsed[0] == bits.CHAT: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                case parsed[0] == bits.BARITONE: return isValidBaritone(args) ? broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1]) : socket.write(responses.BAD_ARGUMENTS)
                case parsed[0] == bits.MOD_COMMAND: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                default: return socket.end(responses.BAD_COMMAND)
            }

    })
    socket.on("error", () => kill(responses.SERVER_ERROR))

}).listen(PORT, HOST)

server.on("connection", (socket) => {
    if (!connectedSockets.has(socket)) connectedSockets.add(socket)
    console.log("New client", socket.remoteAddress)
})

process.on("uncaughtException", () => kill(responses.SERVER_ERROR))
process.on("unhandledRejection", () => kill(responses.SERVER_ERROR))
process.on("SIGINT", () =>  { kill(responses.SERVER_CLOSED); process.exit(0) })


