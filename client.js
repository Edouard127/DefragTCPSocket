const net = require("net")

const client = net.createConnection({
    host: "localhost",
    port: 1984
}, () => {
    console.log("connected to server")
    client.write("4 Kamigen monsupermot2passe")
    client.on("data", (data) => {
        console.log(data.toString().charCodeAt(0))
    })
})