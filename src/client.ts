import net, { Socket } from "net"
import { bits } from "./utils/commands";

const PORT = 1984
const HOST = "localhost"

const args = process.argv.slice(2)

const write = (data: any, socket: Socket) => socket.write(Buffer.from(data).toString("base64")+"\r\n");


const client = net.createConnection(PORT, HOST, () => {
    client.on("data", (data) => {
        const o = Buffer.from(data.toString(), "base64").toString("ascii")

        console.log("Response from server:", o)

        if(o.indexOf("0") === 0) write(bits.HEARTBEAT, client)
    })
    
    client.on("ready", () => {
        process.stdin.on("data", (data) => {
            const stdin = data.toString().trim().split(" ")
            console.log(stdin)
            write(JSON.stringify(stdin), client)
        })
    })
    client.on("close", () => process.exit(0))
    client.on("error", () => process.exit(-1))
})
