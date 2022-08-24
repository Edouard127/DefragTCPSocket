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
        const args = data.toString().split(" ")
        const command = args.shift()
        switch (command) {
            case "16": {
                console.log(data.toString())
            }
        }
    })
})

