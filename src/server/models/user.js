const { userModel } = require('../schemas')
const crypto = require('crypto')

class User {
    constructor ({ username, email, password }) {
        this.approved = [
            username,
            email,
            password
        ].reduce((prev, next) => prev && typeof next !== 'undefined', true)

        this.username = username
        this.email = email
        this.password = password || ''
        this.available = true
    }

    get password () {
        return this._hashedPassword
    }
    set password (val) {
        const hash = crypto.createHash('sha256')
        hash.update(val)
        this._hashedPassword = hash.digest('hex')
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
     * 从数据库中载入该用户
     */
    async load () {
        if (await userModel.exists({ 'username': this.username })) {
            const { username, password, email, available } = await userModel.findOne({ 'username': this.username })
            this.username = username
            this._hashedPassword = password
            this.email = email
            this.available = available
            this.approved = true
        }
    }
}

module.exports = {
    User
}
