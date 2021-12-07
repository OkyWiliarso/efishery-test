const fs = require("fs")

const { generatePass } = require("../helpers/utils")

const create = (payload) => {
    let user = {
        name: payload.name,
        phone: payload.phone,
        password: generatePass(),
        role: payload.role
    }

    fs.readFile("./files/users.json", (err, data) => {
        if (err) return err

        let userList = JSON.parse(data)
        userList.users.push(user)

        fs.writeFileSync("./files/users.json", JSON.stringify(userList))
    })

    return user
}

const checkUserName = (payload) => {
    const data = fs.readFileSync("./files/users.json")

    let userList = JSON.parse(data)
    const checkUser = userList.users.filter((el) => {
        return el.name == payload.name
    })

    if (checkUser.length > 0) {
        return false
    }

    return true
}

module.exports = { create, checkUserName }