const net = require("net")

const client = net.createConnection({
    host: "localhost",
    port: 1984
}, () => {
    console.log("connected to server")
    //client.write(Buffer.from("5 70 1 1 Kamigen password"))
    process.stdin.pipe(client)
    client.on("data", (data) => {
        console.log(data.toString())
    })
})
// Send hello in bytes

