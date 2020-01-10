const mongoose = require('mongoose')
const {
    userAuthorizedKeysModel,
    userModel
} = require('../schemas')
const { User } = require('./user')

class AuthorizedKey {
    constructor ({ _id, user, title, authorizedKey, writed = false }) {
        this._id = _id
        this.user = user
        this.title = title
        this.authorizedKey = authorizedKey
        this.writed = writed
    }

    /**
     *
     * 从数据库中加载authorized_key
     *
     */
    async load () {
        if (typeof this._id === 'undefined') {
            return false
        }
        const { user, authorizedKey, writed } = await userAuthorizedKeysModel.findOne({ '_id': this._id }).populate('user')
        this.user = user
        this.authorizedKey = authorizedKey
        this.writed = writed
    }

    /**
     *
     * 在数据库中创建该authorized_key
     * @return {String} id
     * 
     */
    async create () {
        this._id = new mongoose.Types.ObjectId()
        const key = await userAuthorizedKeysModel.create({
            '_id': this._id,
            'user': this.user._id,
            'authorizedKey': this.authorizedKey,
            'writed': this.writed
        })
        await this.load()
    }

    /**
     *
     * 检查在数据库中是否已经存在该authorized_key
     * @return {Boolean} 是否存在该authorized_key
     *
     */
    async exists () {
        return await userAuthorizedKeysModel.exists({ authorizedKey: this.authorizedKey })
    }
}

module.exports = {
    AuthorizedKey
}
