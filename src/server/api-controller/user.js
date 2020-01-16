const {
    validUserInfo,
    addUser,
    getUserAuthorizedKeys,
    validUserAuthorizedKey,
    createUserAuthorizedKey
} = require('../api-business')
const { User, AuthorizedKey } = require('../models')

/**
 *
 * 获取当前用户身份信息
 * @param {Principal} principal 身份凭证
 * @param {Response} res 响应体
 *
 */
async function getUserInfoAction ({ user }, res) {
    res.json({
        'username': user.username,
        'email': user.email,
        'available': user.available,
        'authorized_keys': await getUserAuthorizedKeys(user)
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
    const entity = {
        'username': req.body.username,
        'email': req.body.email,
        'userPassword': req.body.password,
    }
    const user = new User(entity)
    const validResult = await validUserInfo(entity)
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
