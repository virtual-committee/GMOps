const { userModel } = require('../schemas')

/**
 *
 * 用户注册信息校验
 * @param {User} user 用户注册信息
 * @return {Object} 用户注册信息是否合法
 *
 */
async function validUserInfo (user) {
    if (!user.approved) {
        return {
            'result': false,
            'status': 400,
            'message': {
                'reason': 'mission fields'
            }
        }
    }
    if (!user.valid()) {
        return {
            'result': false,
            'status': 400,
            'message': {
                'reason': 'bad fields'
            }
        }
    }
    if (await user.exists()) {
        return {
            'result': false,
            'status': 409,
            'message': {
                'reason': 'username/email already used'
            }
        }
    }
    return {
        'result': true
    }
}

/**
 *
 * 用户注册
 * @param {User} user 用户实体
 * @return {Boolean} 注册是否成功
 * 
 */
function addUser (user) {
    user.create()
    return true
}

/**
 *
 * 获取用户的全部authorized_keys
 * @param {User} user 用户实体
 * @return {List} authorized_keys
 *
 */
async function getUserAuthorizedKeys (user) {
    return user.getAuthorizedKeys()
}

/**
 *
 * 添加一个authorized_key
 * @param {User} user 用户实体
 * @param {String} title
 * @param {String} authorizedKey
 * @return {String} id
 *
 */
async function createUserAuthorizedKey (user, title, authorizedKey) {
    return await user.createAuthorizedKey(title, authorizedKey)
}

module.exports = {
    validUserInfo,
    addUser,
    getUserAuthorizedKeys,
    createUserAuthorizedKey
}
