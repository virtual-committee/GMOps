const mongoose = require('mongoose')
const {
    userAuthorizedKeysModel,
    userModel
} = require('../schemas')
const { User } = require('./user')

class AuthorizedKey {
    constructor ({ _id = new mongoose.Types.ObjectId(), user, title, authorizedKey, writed = false }) {
        this._id = _id
        this.user = user
        this.title = title
        this.authorizedKey = authorizedKey
        this.writed = writed
        this.approved = !!authorizedKey
    }

    _sync ({ user, title, authorizedKey, writed }) {
        this.user = user
        this.title = title
        this.authorizedKey = authorizedKey
        this.writed = writed
        this.approved = true
    }

    /**
     *
     * 从数据库中加载authorized_key
     * @return {Boolean} 加载是否成功
     *
     */
    async load () {
        this.approved = false
        if (typeof this._id === 'undefined') {
            return false
        }
        if (!await userAuthorizedKeysModel.exists({ _id: this._id })) {
            return false
        }
        this._sync(await userAuthorizedKeysModel.findOne({ _id: this._id }).populate('user'))
        return true
    }

    /**
     *
     * 在数据库中创建该authorized_key
     * @return {String} id
     * 
     */
    async create () {
        const key = await userAuthorizedKeysModel.create({
            _id: this._id,
            user: this.user._id,
            authorizedKey: this.authorizedKey,
            writed: this.writed
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

    /**
     *
     * 标记为写入状态
     *
     */
    async markWrite () {
        let obj = await userAuthorizedKeysModel.findOne({ _id: this._id })
        obj.writed = true
        this._sync(await obj.save())
    }
}

module.exports = {
    AuthorizedKey
}
