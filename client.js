import net from "net"
import { bits } from "./utils/commands.js"
import zlib from "node:zlib"
import crypto from "crypto"

const PORT = 1984
const HOST = "localhost"
const pass = crypto.randomUUID()

const client = net.connect(PORT, HOST, () => {
        client.on("ready", () => {
            console.log("Connected, password:", pass)
            console.log(JSON.stringify([0, pass]))
            client.write(JSON.stringify([0, pass])) // 0 is init with password
        })
        client.on("data", (data) => {
            console.log(data.toString())
        })  
})