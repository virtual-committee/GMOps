const log4js = require('log4js')

module.exports = function (app) {
    app.use(log4js.connectLogger(log4js.getLogger('api'), {
        format: 'GMOps API server - :method :url :status'
    }))

    require('./user').map(route => app.use('/user', route))
    require('./authorized-key').map(route => app.use('/authorized-key', route))
}
