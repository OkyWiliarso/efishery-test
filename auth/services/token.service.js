const jwt = require("jsonwebtoken")

const accessToken = (user) => {
    const token = jwt.sign({
        name: user.name,
        role: user.role,
        phone: user.phone,
        timestamp: user.created_at
    },
    "secretkey",
    { expiresIn: '1d', algorithm: 'HS256' }
    )
  
    return { token }
}

module.exports = { accessToken }