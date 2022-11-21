const net = require("net")
require("colors")

const client = net.createConnection({
    host: "localhost",
    port: 1984
}, () => {
    console.log("Connected to server".green)
    let byteArray = [18, 1, 135,115,247,39,59,137,87,88,4,36,238,175,125,57,80,3]
    client.write(Buffer.from(byteArray))
    process.stdin.pipe(client)
    client.on("data", (data) => {
        console.log(data.toString()+"\n\n")
    })
})

