const {
    userModel,
    userAuthorizedKeysModel,
    userRepoModel
} = require('../schemas')
const crypto = require('crypto')
const hash = crypto.createHash('sha256')
const { AuthorizedKey } = require('./authorized-key')
const { Repo } = require('./repo')
const mongoose = require('mongoose')

class User {
    constructor ({ _id = new mongoose.Types.ObjectId(), username, email, userPassword, password, available = true }) {
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

    _sync ({ _id, username, password, email, available }) {
        this._id = _id
        this.username = username
        this._hashedPassword = password
        this.email = email
        this.available = available
        this.approved = true
    }

    valid () {
        // TODO 对 username/email/password进行验证
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
            $or: [
                { username: this.username },
                { email: this.email }
            ]
        })
    }

    /**
     *
     * 在数据库中创建该用户
     */
    create () {
        return userModel.create({
            _id: this._id,
            username: this.username,
            password: this.password,
            email: this.email,
            available: this.available
        })
    }

    /**
     *
     * 从数据库中载入该用户
     *
     */
    async load () {
        this.approved = false
        if (await userModel.exists({ username: this.username })) {
            this._sync(await userModel.findOne({ username: this.username }))
        }
    }

    /**
     *
     * 获取当前用户下的全部authorized_keys
     * @return {List} authorized_keys
     *
     */
    async getAuthorizedKeys () {
        return await userAuthorizedKeysModel.
            find({ user: this._id }).
            map(res => res.map(item => new AuthorizedKey(item)))
    }

    /**
     *
     * 获取当前用户下的全部Repo
     * @return {List} Repos
     *
     */
    async getRepos () {
        return await userRepoModel.
            find({ user: this._id }).
            populate('repo').
            map(res => res.map(({ repo }) => {
                const ret = new Repo(repo)
                ret.user = this
                return ret
            }))
    }
}

module.exports = {
    User
}
