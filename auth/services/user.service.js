const fs = require("fs")

const { generatePass } = require("../helpers/utils")


const create = (payload) => {
    let user = {
        name: payload.name,
        phone: payload.phone,
        password: generatePass(),
        role: payload.role
    }

    fs.readFile("./files/users.json", function readFileCallback(err, data){
        if (err){
            return err
        } else {
            let userList = JSON.parse(data)
            userList.users.push(user)

            fs.writeFileSync("./files/users.json", JSON.stringify(userList))
        }
    })
   
    return user
}

module.exports = { create }