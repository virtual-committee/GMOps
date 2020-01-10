const {
    validUserInfo,
    addUser
} = require('../business/user')
const { User } = require('../models')

/**
 *
 * 获取当前用户身份信息
 * @param {Principal} principal 身份凭证
 * @param {Response} res 响应体
 *
 */
async function getUserInfoAction ({ principal }, res) {
    const user = new User(principal)
    await user.load()
    if (!user.approved) {
        res.status(404).json({ 'reason': 'the user dose not exist' }).end()
    }
    else {
        res.json({
            'username': user.username,
            'email': user.email,
            'available': user.available
        }).end()
    }
}

/**
 *
 * 用户注册
 * @param {Request} req
 * @param {Response} res
 *
 */
async function userRegisterAction (req, res) {
    const user = new User(req.body)
    const validResult = await validUserInfo(user)
    if (!validResult.result) {
        const { status, message } = validResult
        res.status(status).json(message).end()
        return
    }
    if (!addUser(user)) {
        res.status(500).json({ 'status': 'Server Internal Error' }).end()
        return
    }
    res.status(201).json({ 'status': 'created' }).end()
}

module.exports = {
    getUserInfoAction,
    userRegisterAction
}
