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

module.exports = {
    validUserInfo,
    addUser,
}
