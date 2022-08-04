import net from "net"
import { bits } from "./utils/commands.js"
import zlib from "node:zlib"
import crypto from "crypto"

const PORT = 1984
const HOST = "0.0.0.0"

const connectedSockets = new Set();

let pass = ""

const broadcast = function (data, sender, password) {
    console.log(data)
    console.log(pass)

    for (let sock of connectedSockets) {
        if (sock !== sender && crypto.createHash('sha256').update(password).digest('base64') == pass) {
            sock.write(JSON.stringify(data)+"\r\n");
        }
    }
}

const server = net.createServer((socket) => {
    socket.on("data", (data) => {
            console.log(data.toString())

            const command = /*zlib.inflateSync(data.buffer).toString()*/data.toString()

            if (!command) return socket.end("-1")

            const parsed = [...JSON.parse(command)]

            const args = [...parsed].slice(2)
            console.log(parsed, args)
            
            switch(true) {
                case parsed[0] == bits.INIT && !pass: return pass = crypto.createHash('sha256').update(parsed[1]).digest('base64')
                case parsed[0] == bits.CONNECT: return broadcast([parsed[0], ...parsed.splice(2, 2)], socket, parsed[1])
                case parsed[0] == bits.CHAT: return broadcast([parsed[0], ...parsed.splice(2)], socket, parsed[1])
            }

    })
    socket.on("error", (err) => {
        console.log(err.message)
    })

}).listen(PORT, HOST)

server.on("connection", (socket) => {
    if (!connectedSockets.has(socket)) connectedSockets.add(socket)
    console.log("New client", socket.remoteAddress)
})
server.on("close", () => connectedSockets.clear())

