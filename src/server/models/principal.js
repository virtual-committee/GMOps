const { userModel } = require('../schemas/user')

class Principal {
	constructor (username) {
		this.username = username
		this.exists = false
		this.synced = false
	}

	async syncDB () {
		if (this.synced) {
			return
		}
		const user = await userModel.findOne({ username: this.username })
		this.exists = !!user
		if (this.exists) {
			this.email = user.email
			this.available = user.available
		}
		this.synced = false
	}
}

module.exports = {
	Principal
}
