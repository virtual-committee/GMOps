const express = require('express')
const fs = require('fs')
const log4js = require('log4js')
const apiRouteStuffing = require('./routes/api')
const webRouteStuffing = require('./routes/web')

const api = express()
const web = express()

log4js.configure('./src/server/config/log4js.json')
const logger = log4js.getLogger('system')

apiRouteStuffing(api)
webRouteStuffing(web)

logger.info('GMOps web server listening at \'http://0.0.0.0:8080\'')
web.listen(8080)

const apiUDS = '/var/run/gmops.sock'
fs.exists(apiUDS, function (exists) {
    if (exists) {
        fs.unlinkSync(apiUDS)
    }
    logger.info('GMOps API server listening at \'%s\'', apiUDS)
    api.listen(apiUDS)
})
