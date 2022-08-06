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
    })
    
    client.on("ready", () => {
        process.stdin.on("data", (data) => {
            const command = data.toString()
            console.log("Command:", command)
            client.write(command)
        })
    })
    client.on("close", () => process.exit(0))
    client.on("error", () => process.exit(-1))
})
