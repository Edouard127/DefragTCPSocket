const net = require("net")
require("colors")

const client = net.createConnection({
    host: "localhost",
    port: 1984
}, () => {
    console.log("Connected to server".green)
    client.write(Buffer.from("13 0"))
    process.stdin.pipe(client)
    client.on("data", (data) => {
        console.log(data.toString())
    })
})

