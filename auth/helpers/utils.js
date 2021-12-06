const responseSuccess = (res, obj) => {
    res
        .status(200)
        .send({
            data: obj.data || {}
        })
  }
  
const responseError = (res, obj) => {
    res
        .status(obj.status || 501)
        .send({
            error: obj.message
        })
}

const validatePhone = (phoneNumber) => {
    const phoneNum = phoneNumber.replace(/[^\d]/g, '')
    if (phoneNum.length > 6 && phoneNum.length < 13 && phoneNum[0] == 0) return true
}

const generatePass = () => {
    const size = 4
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()"
    let result = ""

    for(let i = 0; i < size; i++) {
        result += charset[Math.floor(Math.random() * charset.length)]
    }

    return result
}
  
module.exports = { responseSuccess, responseError, validatePhone, generatePass}