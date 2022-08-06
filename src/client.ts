import net, { Socket } from "net"
import { bits } from "./utils/commands.js"
import zlib from "node:zlib"

const PORT = 1984
const HOST = "localhost"

const args = process.argv.slice(2)

const write = (data: any, socket: Socket) => socket.write(zlib.deflateSync(data+"\r\n"));


const client = net.createConnection(PORT, HOST, () => {
    client.on("data", (data) => {
        console.log(data.toString())
        const o = zlib.inflateSync(data.buffer).toString().trim()
        console.log("Response from server:", o)

        const response = o.replace("PING", "0").split(" ")
        console.log("Response:", response)
        if(o.indexOf("PING") > -1) write(JSON.stringify(response), client)
    })
    
    client.on("ready", () => {
        process.stdin.on("data", (data) => {
            const stdin = data.toString().trim().split(" ")
            console.log(JSON.stringify(stdin))
            write(JSON.stringify(stdin), client)
        })
    })
    client.on("close", () => process.exit(0))
    client.on("error", () => process.exit(-1))
})
