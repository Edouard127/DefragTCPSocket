const net = require("net")
require("colors")

const client = net.createConnection({
    host: "localhost",
    port: 1984
}, () => {
    console.log("Connected to server".green)
    client.write(Buffer.from("10 0 13 0"))
    process.stdin.pipe(client)
    client.on("data", (data) => {
        console.log(data.toString()+"\n\n")
        //const args = data.toString().split(" ")
        //const command = args.shift()
        /*switch (command) {
            case "16": {

            }
        }*/
    })
})

