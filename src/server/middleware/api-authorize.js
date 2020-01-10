const { apiAuthorize } = require('../business')
const {
    validPrincipal,
    buildPrincipal
} = require('../business')

/**
 *
 * 凭证中间件，用于构造凭证并验证凭证是否合法
 * @param {Request} req
 * @param {Response} res
 * @param {Function} next
 *
 */
async function authorize (req, res, next) {
    req.principal = buildPrincipal(req.get('GMOps-Username'))
    if (await validPrincipal(req.principal)) {
        next()
    }
    else {
        res.status(401).json({ 'status': 'Unauthorized' })
    }
}

module.exports = {
    authorize
}
