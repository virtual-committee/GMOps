const { Repo } = require('../models')

/**
 *
 * 获取所属用于的全部Repos
 * @param {User} user
 * @return {List} Repos
 *
 */
async function getUserRepos (user) {
    return (await user.getRepos()).map(repo => ({
        'name': repo.name,
        'descript': repo.descript,
        'attr': repo.attr,
        'user': user.username
    }))
}

/**
 *
 * 查看当前用户是否存在指定的Repo
 * @param {User} user
 * @param {String} 仓库名
 * @return {Boolean} 是否存在该仓库名的Repo
 *
 */
async function existsUserRepo (user, repoName) {
    return !!(await user.getRepos()).find(repo => repo.name === repoName)
}

/**
 *
 * 创建用户仓库
 * @param {User} user
 * @param {String} repoName 用户
 * @param {String} descript 仓库描述
 * @return {Object}
 *
 */
async function createUserRepo (user, repoName, descript) {
    if (await existsUserRepo(user, repoName)) {
        return {
            'result': false,
            'status': 409,
            'message': {
                'reason': 'repo already exists'
            }
        }
    }
    const repo = new Repo({
        user,
        name: repoName,
        descript,
        attr: 0o711
    })
    if (!await repo.create()) {
        return {
            'result': false,
            'status': 500,
            'message': {
                'reason': 'Server Internal Error'
            }
        }
    }
    return {
        'result': true
    }
}

module.exports = {
    getUserRepos,
    createUserRepo
}
