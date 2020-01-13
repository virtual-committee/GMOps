const { AuthorizedKey } = require('../models')
const {
    validAuthorizedKey,
    applyAuthorizedKey
} = require('../business')

async function applyAuthorizedKeyAction (req, res) {
    const { id } = req.params
    const authorizedKey = new AuthorizedKey({ '_id': id })
    const ret = await validAuthorizedKey(authorizedKey)
    if (!ret.result) {
        res.status(ret.status).json(ret.message).end()
        return
    }
    if (!applyAuthorizedKey(authorizedKey)) {
        res.status(500).json({ 'reason': 'Server Internal Error' }).end()
        return
    }
    authorizedKey.writed = true
    authorizedKey.markWrite()
    res.status(201).json({ 'status': 'applied' }).end()
}

module.exports = {
    applyAuthorizedKeyAction
}
