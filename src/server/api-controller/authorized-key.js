const { AuthorizedKey } = require('../models')
const {
    validAuthorizedKey,
    applyAuthorizedKey,
    cancelAuthorizedKey,
    fastCancelAuthorizedKey
} = require('../business')

/**
 *
 * 将authorized_key写入到authorized_keys文件中
 * @param {String} id
 * @param {Response} res
 *
 */
async function applyAuthorizedKeyAction ({ params: { id } }, res) {
    const authorizedKey = new AuthorizedKey({ '_id': id })
    const ret = await validAuthorizedKey(authorizedKey)
    if (!ret.result) {
        res.status(ret.status).json(ret.message).end()
        return
    }
    if (!await applyAuthorizedKey(authorizedKey)) {
        res.status(500).json({ 'reason': 'Server Internal Error' }).end()
        return
    }
    res.status(201).json({ 'status': 'applied' }).end()
}

/**
 *
 * 将authorized_key从authorized_keys文件中删除
 * @param {String} id
 * @param {Response} res
 *
 */
async function cancelAuthorizedKeyAction ({ params: { id } }, res) {
    const authorizedKey = new AuthorizedKey({ '_id': id })
    const cancelRet = await cancelAuthorizedKey(authorizedKey)
    if (!cancelRet.result) {
        const { status, message } = cancelRet
        res.status(status).json(message).end()
        return
    }
    res.status(202).json({ 'status': 'canceled' }).end()
}

/**
 *
 * 快速撤销authorized_key
 * @param {String} id
 * @param {Response} res
 *
 */
async function fastCancelAuthorizedKeyAction ({ params: { id } }, res) {
    const authorizedKey = new AuthorizedKey({ '_id': id })
    const cancelRet = await fastCancelAuthorizedKey(authorizedKey)
    if (!cancelRet.result) {
        const { status, message } = cancelRet
        res.status(status).json(message).end()
        return
    }
    res.status(202).json({ 'status': 'canceled' }).end()
}

module.exports = {
    applyAuthorizedKeyAction,
    cancelAuthorizedKeyAction,
    fastCancelAuthorizedKeyAction
}
