const { responseError, responseSuccess, validatePhone } = require("../helpers/utils")
const { getUser } = require("../services/user.service")
const { accessToken, validateToken } = require("../services/token.service")

const generateToken = (req, res) => {
    try {
        if (!validatePhone(req.body.phone)) throw { status: 400, message: "Invalid PhoneNumber" }
        
        const user = getUser(req.body)
        if (!user) throw { status: 400, message: "Wrong PhoneNumber / Password" }

        const result = accessToken(user)
        return responseSuccess(res, {
            data: result
        })
    } catch (err) {
        return responseError(res, err)
    }
}

const verifyToken = (req, res) => {
    try {
        const result = validateToken(req.body)
        return responseSuccess(res, {
            data: result
        })
    } catch (err) {
        return responseError(res, err)
    }
}

module.exports = { generateToken, verifyToken }