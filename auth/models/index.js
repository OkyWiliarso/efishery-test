const { DB } = require("../../config.json")
const Sequelize = require("sequelize")

const sequelize = new Sequelize(DB.DB_NAME, DB.USER, DB.PASSWORD, {
  host: DB.HOST,
  dialect: DB.dialect,
  operatorsAliases: false,

  pool: {
    max: DB.pool.max,
    min: DB.pool.min,
    acquire: DB.pool.acquire,
    idle: DB.pool.idle
  }
})

const db = {}

db.Sequelize = Sequelize
db.sequelize = sequelize

db.user = require("./user")(sequelize, Sequelize)


module.exports = db