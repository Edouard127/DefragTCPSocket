import net, { Socket } from "net"
import { bits } from "./utils/commands.js"
import { isValidBaritone } from "./utils/isValidBaritone"
import zlib from "node:zlib"
import crypto from "crypto"

const PORT = 1984
const HOST = "0.0.0.0"

const connectedSockets = new Set<Socket>();

let pass = ""

const goodPass = (password: string) => crypto.createHash('sha256').update(password).digest('base64') == pass

const broadcast = (data: Array<string>, sender: Socket, password: string) => {
    connectedSockets.forEach((sock) => {
        sock = sock as Socket
        if (sock !== sender && goodPass(password)) {
            sock.setEncoding('utf8');
            sock.write(data.toString().replace(/,/g, " ")+"\r\n");
        }
    })
}

const kill = () => connectedSockets.forEach((sock) => sock.end("-1"))

const server = net.createServer((socket) => {
    socket.on("data", (data) => {
            console.log(data.toString())

            const command = /*zlib.inflateSync(data.buffer).toString()*/data.toString()

            if (!command) return socket.end("-1")

            const parsed = [...JSON.parse(command)]

            const args = [...parsed].slice(2)
            console.log(parsed, args)
            
            switch(true) {
                case parsed[0] == bits.EXIT && goodPass(parsed[1]): pass = ""; return kill()
                case parsed[0] == bits.INIT && !pass: return pass = crypto.createHash('sha256').update(parsed[1]).digest('base64')
                case parsed[0] == bits.CONNECT: return broadcast([parsed[0], ...parsed.splice(2, 2)], socket, parsed[1])
                case parsed[0] == bits.CHAT: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                case parsed[0] == bits.BARITONE: if(isValidBaritone(args)) return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                case parsed[0] == bits.MOD_COMMAND: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
                default: return socket.end("-1")
            }

    })
    socket.on("error", () => kill())

}).listen(PORT, HOST)

server.on("connection", (socket) => {
    if (!connectedSockets.has(socket)) connectedSockets.add(socket)
    console.log("New client", socket.remoteAddress)
})

