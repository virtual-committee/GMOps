const log4js = require('log4js')

module.exports = function (app) {
    app.use(log4js.connectLogger(log4js.getLogger('web'), {
        format: 'GMOps web server - :method :url :status'
    }))
}
