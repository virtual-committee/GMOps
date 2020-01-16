const {
    getUserRepos,
    createUserRepo
} = require('../api-business')

/**
 *
 * 获取当前用户的所有Repos
 * @param {User} user 用户
 * @param {Response} res 响应体
 *
 */
async function getUserReposAction ({ user }, res) {
    res.json(await getUserRepos(user)).end()
}

/**
 *
 * 当前用户创建Repo
 * @param {User} user 用户
 * @param {String} descript
 * @param {String} repoName 仓库名
 * @param {Response} res 响应体
 *
 */
async function createUserRepoAction ({ user, body: { descript }, params: { repoName } }, res) {
    const result = await createUserRepo(user, repoName, descript)
    if (!result.result) {
        const { status, message } = result
        res.status(status).json(message).end()
        return
    }
    res.status(201).json({ 'status': 'repo created' }).end()
}

module.exports = {
    getUserReposAction,
    createUserRepoAction
}
