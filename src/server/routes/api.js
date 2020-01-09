const log4js = require('log4js')
const authorize = require('../middleware/api-authorize')

log4js.configure('./src/server/config/log4js.json')

module.exports = function (app) {
    app.use(log4js.connectLogger(log4js.getLogger('api'), {
        format: 'GMOps API server - :method :url :status'
    }))
    .use(authorize)
}
