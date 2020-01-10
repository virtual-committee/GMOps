const mongoose = require('mongoose')

const userSchema = new mongoose.Schema({
    _id: mongoose.Schema.Types.ObjectId,
    username: { type: String, required: true, unique: true },
    password: String,
    email: { type: String, required: true, unique: true },
    available: Boolean
})


const userAuthorizedKeysSchema = new mongoose.Schema({
    _id: mongoose.Schema.Types.ObjectId,
    user: { type: mongoose.Schema.Types.ObjectId, ref: 'User', unique: false },
    title: { type: String, require: true },
    authorizedKey: { type: String, required: true, unique: true },
    writed: { type: Boolean, required: true }
})


module.exports = {
    userModel: mongoose.model('User', userSchema),
    userAuthorizedKeysModel: mongoose.model('UserAuthorizedKeys',userAuthorizedKeysSchema)
}
