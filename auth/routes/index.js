const router = require("express").Router()

router.get("/ping", (_req, res) => res.status(200).json({ message: "Pong!" }))

module.exports = router
