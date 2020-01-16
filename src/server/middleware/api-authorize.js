const {
    apiAuthorize,
    validPrincipal,
    buildPrincipal
} = require('../api-business')
const { User } = require('../models')

/**
 *
 * 凭证中间件，用于构造凭证并验证凭证是否合法
 * @param {Request} req
 * @param {Response} res
 * @param {Function} next
 *
 */
async function authorize (req, res, next) {
    const principal = buildPrincipal(req.get('GMOps-Username'))
    if (await validPrincipal(principal)) {
        req.user = new User(principal)
        await req.user.load()
        if (!req.user.approved) {
            res.status(404).json({ 'reason': 'the user dose not exist' }).end()
        }
        else {
            next()
        }
    }
    else {
        res.status(401).json({ 'status': 'Unauthorized' })
    }
}

module.exports = {
    authorize
}
