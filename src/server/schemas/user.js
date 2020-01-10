const mongoose = require('mongoose')

const userSchema = new mongoose.Schema({
    username: { type: String, required: true, unique: true },
    password: String,
    email: { type: String, required: true, unique: true },
    available: Boolean
})

module.exports = {
	userModel: mongoose.model('User', userSchema)
}
