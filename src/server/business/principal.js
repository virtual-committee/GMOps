const { Principal } = require('../models/principal')

/**
 *
 * 验证凭证
 * @param {Principal} 凭证
 * @return {Boolean} 该凭证是否合法
 *
 */
async function validPrincipal (principal) {
	await principal.syncDB()
	if (!principal.exists) {
		console.log('here')
		return false
	}
	return principal.available
}

/**
 *
 * 根据传递到服务端的Username构造凭证
 * @param {String} username 用户名
 * @return {Principal} 用户凭证
 *
 */
function buildPrincipal (username) {
	return new Principal(username)
}

module.exports = {
	validPrincipal,
	buildPrincipal
}
