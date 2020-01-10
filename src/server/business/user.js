const { userModel } = require('../schemas/user')

/**
 *
 * 用户注册信息校验
 * @param {Object} info 用户注册信息
 * @return {Object} 用户注册信息是否合法
 *
 */
async function validUserInfo (info) {
	if (typeof info.username === 'undefined') {
		return {
			'result': false,
			'status': 400,
			'message': {
				'reason': 'missing \'username\'' 
			}
		}
	}
	if (typeof info.password === 'undefined') {
		return {
			'result': false,
			'status': 400,
			'message': {
				'reason': 'mission \'password\''
			}
		}
	}
	if (typeof info.email === 'undefined') {
		return {
			'result': false,
			'status': 400,
			'message': {
				'reason': 'mission \'email\''
			}
		}
	}
	// TODO 对用户名、密码和电子邮件格式的检查
	if (await userModel.exists({
		$or: [
			{ username: info.username },
			{ email: info.email }
		]
	})) {
		return {
			'result': false,
			'status': 409,
			'message': {
				'reason': 'username/email already used'
			}
		}
	}
	return true
}

/**
 *
 * 用户注册
 * @param {info} info 用户注册凭证
 * @return {Boolean} 注册是否成功
 * 
 */
function addUser ({ username, email, password }) {
	userModel.create({
		username,
		password,
		email,
		available: true
	})
	return true
}

module.exports = {
	validUserInfo,
	addUser
}
