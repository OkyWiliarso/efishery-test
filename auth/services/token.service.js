const jwt = require("jsonwebtoken")

const { TOKEN } = require("../config.json")

const accessToken = (user) => {
    const token = jwt.sign({
        name: user.name,
        role: user.role,
        phone: user.phone,
        timestamp: user.created_at
    },
    TOKEN.SECRET,
    { expiresIn: '1d', algorithm: 'HS256' }
    )
  
    return { token }
}

const validateToken = (payload) => {
    const tokenPayload = jwt.verify(payload.token, TOKEN.SECRET, (err, payload) => {
        if (err) throw { status: 401, message: 'Invalid Token' }
    
        return payload
      })
    
      return tokenPayload
}

module.exports = { accessToken, validateToken }