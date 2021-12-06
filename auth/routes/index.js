const router = require("express").Router()
const { createUser } = require("../controllers/user.controller")

router.get("/ping", (_req, res) => res.status(200).json({ message: "Pong!" }))
router.post("/user", createUser)

module.exports = router
