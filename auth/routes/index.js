const router = require("express").Router()
const { createUser } = require("../controllers/user.controller")
const { generateToken } = require("../controllers/token.controller")

router.get("/ping", (_req, res) => res.status(200).json({ message: "Pong!" }))
router.post("/user", createUser)
router.post("/token/create", generateToken)

module.exports = router
