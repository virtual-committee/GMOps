const { userModel } = require('../schemas/user')

class Principal {
    constructor (username) {
        this.username = username
        this.available = false
        this.exists = false
        this.synced = false
    }

    async syncDB () {
        if (this.synced) {
            return
        }
        const userDoc = await userModel.findOne({ 'username': this.username })
        this.exists = !!userDoc
        this.available = userDoc.available
        this.synced = true
    }
}

module.exports = {
    Principal
}
