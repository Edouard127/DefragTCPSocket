const net = require("net")

const client = net.createConnection({
    host: "localhost",
    port: 1984
}, () => {
    console.log("connected to server")
    client.write("5 Kamigen password")
    client.on("data", (data) => {
        console.log(data.toString().charCodeAt(0))
    })
})