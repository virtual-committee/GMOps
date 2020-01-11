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
 * @param {AuthorizdKey} authorizedKey
 * @return {List} authorized_keys
 *
 */
async function getUserAuthorizedKeys (user) {
    return user.getAuthorizedKeys()
}

/**
 *
 * 验证AuthorizedKey合法性
 * @param {AuthorizedKey} authorizedKey
 * @result {Object} 
 *
 */
async function validUserAuthorizedKey (authorizedKey) {
    if (!authorizedKey.user.approved) {
        return {
            'result': false,
            'status': 404,
            'message': {
                'reason': 'user dose not exist'
            }
        }
    }
    if (await authorizedKey.exists()) {
        return {
            'result': false,
            'status': 409,
            'message': {
                'reason': 'authorized_key already used'
            }
        }
    }
    return {
        'result': true
    }
}

/**
 *
 * 添加一个authorized_key
 * @param {AuthorizeKey} authorizedKey
 * @return {Object} 执行状态
 *
 */
async function createUserAuthorizedKey (authorizedKey) {
    await authorizedKey.create()
    return {
        'status': 201,
        'message': {
            'id': authorizedKey._id
        }
    }
}

module.exports = {
    validUserInfo,
    addUser,
    getUserAuthorizedKeys,
    validUserAuthorizedKey,
    createUserAuthorizedKey
}
