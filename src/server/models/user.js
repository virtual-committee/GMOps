const {
    userModel,
    userAuthorizedKeysModel
} = require('../schemas')
const crypto = require('crypto')
const hash = crypto.createHash('sha256')
const { AuthorizedKey } = require('./authorized-key')
const mongoose = require('mongoose')

class User {
    constructor ({ _id, username, email, userPassword, password, available = true }) {
        this.approved = [
            username,
            email,
            userPassword
        ].reduce((prev, next) => prev && typeof next !== 'undefined', true)

        this._id = _id
        this.username = username
        this.email = email
        this.userPassword = userPassword
        if (typeof userPassword !== 'undefined') {
            this.password = userPassword
        }
        else if (typeof password !== 'undefined') {
            this._hashedPassword = password
        }
        else {
            this.password = ''
        }
        this.available = available
    }

    get password () {
        return this._hashedPassword
    }
    set password (val) {
        if (val === '') {
            this._hashedPassword = ''
        }
        else {
            hash.update(val)
            this._hashedPassword = hash.digest('hex')
        }
    }

    valid () {
        // 对 username/email/password进行验证
        return true
    }

    /**
     *
     * 验证该用户是否存在在数据库中
     *
     */
    async exists () {
        if (!this.approved) {
            return false
        }
        return await userModel.exists({
            '$or': [
                { 'username': this.username },
                { 'email': this.email }
            ]
        })
    }

    /**
     *
     * 在数据库中创建该用户
     */
    create () {
        return userModel.create({
            'username': this.username,
            'password': this.password,
            'email': this.email,
            'available': this.available
        })
    }

    /**
     *
     * 从数据库中载入该用户
     *
     */
    async load () {
        if (await userModel.exists({ 'username': this.username })) {
            const { _id, username, password, email, available } = await userModel.findOne({ 'username': this.username })
            this._id = _id
            this.username = username
            this._hashedPassword = password
            this.email = email
            this.available = available
            this.approved = true
        }
    }

    /**
     *
     * 获取当前用户下的全部authorized_keys
     *
     */
    async getAuthorizedKeys () {
        return await userAuthorizedKeysModel.
            find({ 'user': this._id }).
            populate('user').
            map(res => res.map(item => new AuthorizedKey(item)))
    }

    /**
     *
     * 添加一个新的authorized_key
     * @param {String} title authorized_key标识
     * @param {String} authorizedKey authorized_key
     * @return {String} 创建成功后 authorized_key的ID
     *
     */
    async createAuthorizedKey (title, authorizedKey) {
        const key = new AuthorizedKey({
            'user': this,
            title,
            authorizedKey
        })
        await key.create()
        return key._id
    }
}

module.exports = {
    User
}
