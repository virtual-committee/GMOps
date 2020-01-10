const { apiAuthorize } = require('../business')


module.exports = function (req, res, next) {
    if (!apiAuthorize(req)) {
        res.status(401).send('Unauthorized')
        return
    }
    next()
}
