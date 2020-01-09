module.exports = function (req, res, next) {
    if (typeof req.get('GMOps-Username') === 'undefined') {
        res.status(401).send('Unauthorized')
    }
    next()
}
