const {
	validUserInfo,
	addUser
} = require('../business/user')

/**
 *
 * 获取当前用户身份信息
 * @param {Principal} principal 身份凭证
 * @param {Response} res 响应体
 *
 */
function userInfoAction ({ principal }, res) {
	res.send({
		'username': principal.username
	}).end()
}

/**
 *
 * 用户注册
 * @param {Request} req
 * @param {Response} res
 *
 */
async function userRegisterAction (req, res) {
	const validResult = await validUserInfo(req.body)
	if (!validResult.result) {
		const { status, message } = validResult
		res.status(status).json(message).end()
	}
	// TODO 对password进行加密/哈希处理
	if (!addUser(req.body)) {
		res.status(500).json({ 'status': 'Server Internal Error' }).end()
		return
	}
	res.status(201).json({ 'status': 'created' }).end()
}

module.exports = {
	userInfoAction,
	userRegisterAction
}
