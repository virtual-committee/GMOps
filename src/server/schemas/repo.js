const mongoose = require('mongoose')

const repoSchema = new mongoose.Schema({
    owner: { type: String, required: true },
    name: { type: String, required: true },
    realPath: { type: String, required: true, unique: true }
})

module.exports = {
    repoSchema
}
