const express = require("express")
const http = require("http")
const logger = require("morgan")
const routes = require("./routes")
const { PORT } = require('../config.json')

const app = express()

app.use(logger("dev"))

app.use("/", routes)

const server = http.createServer(app)
server.listen(PORT.API, () => console.log(`Listen to the port of ${PORT.API}`))