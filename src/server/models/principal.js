const mongoose = require('mongoose')
const { userModel } = require('../schemas/user')

class Principal {
	constructor (username) {
		this.username = username
		this.loaded = false
		this.available = false
	}

	async valid () {
		if (this.loaded) {
			return this.available
		}
		if (typeof this.username === 'undefined') {
			return false
		}
		const res = await userModel.where({ username: this.username }).findOne().map(res => res)
		this.loaded = true
		if (!res) {
			return false
		}
		if (!res.available) {
			return false
		}
		return true
	}
}

module.exports = {
	Principal
}
