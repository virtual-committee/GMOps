const mongoose = require('mongoose')

const repoSchema = new mongoose.Schema({
    name: { type: String, required: true },
    realPath: { type: String, required: true, unique: true }
})

module.exports = {
    repoSchema
}
