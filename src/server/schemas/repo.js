const mongoose = require('mongoose')

const repoSchema = new mongoose.Schema({
    name: { type: String, required: true },
    descript: { type: String, required: false },
    attr: { type: Number, required: true }
})

const userRepoSchema = new mongoose.Schema({
    user: { type: mongoose.Schema.Types.ObjectId, ref: 'User', unique: false },
    repo: { type: mongoose.Schema.Types.ObjectId, ref: 'Repo', unique: true }
})

module.exports = {
    repoModel: mongoose.model('Repo', repoSchema),
    userRepoModel: mongoose.model('UserRepo', userRepoSchema)
}
