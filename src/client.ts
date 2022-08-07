import net, { Socket } from "net"

const PORT = 1984
const HOST = "localhost"

const args = process.argv.slice(2)

const write = (data: any, socket: Socket) => socket.write(Buffer.from(data).toString("base64")+"\r\n");


const client = net.createConnection(PORT, HOST, () => {
    client.on("data", (data) => {
        const o = Buffer.from(data.toString(), "base64").toString("ascii")

        console.log("Response from server:", o)

        const response = o.split(" ")

        if(o.indexOf("0") === 0) write(JSON.stringify(response), client)
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
