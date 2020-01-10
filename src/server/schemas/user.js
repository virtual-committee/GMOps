const mongoose = require('mongoose')

const userSchema = new mongoose.Schema({
    username: { type: String, required: true, unique: true },
    password: String,
    email: { type: String, required: true, unique: true },
    available: Boolean
})


const userAuthorizedKeysSchema = new mongoose.Schema({
    user: { type: mongoose.ObjectId, ref: 'User' },
    authorizedKey: { type: String, required: true, unique: true },
    writed: { type: Boolean, required: true }
})


module.exports = {
    userModel: mongoose.model('User', userSchema),
    userAuthorizedKeysModel: mongoose.model('UserAuthorizedKeys',userAuthorizedKeysSchema)
}
