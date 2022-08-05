import net from "net"
import { bits } from "./utils/commands.js"
import zlib from "node:zlib"

const PORT = 1984
const HOST = "localhost"

const args = process.argv.slice(2)
console.log(args)

const client = net.connect(PORT, HOST, () => {
    client.on("data", (data) => {
        console.log("Response from server:", data.toString())
        client.end()
    })
    
    client.on("ready", () => {
        const command = /*zlib.deflateSync(*/JSON.stringify([...args])//)
        client.write(command) // Send
    })
})