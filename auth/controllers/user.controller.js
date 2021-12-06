const { responseError, responseSuccess, validatePhone } = require("../helpers/utils")
const { create } = require("../services/user.service")

const createUser = (req, res) => {
    try {
        if (!validatePhone(req.body.phone)) throw { status: 400, message: "Invalid Phonenumber" }

        const result = create(req.body)
        return responseSuccess(res, {
            data: result
        })
    } catch (err) {
        return responseError(res, err)
    }
}

module.exports = { createUser }