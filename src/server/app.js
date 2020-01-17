const express = require('express')
const fs = require('fs')
const log4js = require('log4js')
const apiRouteStuffing = require('./routes/api')
const webRouteStuffing = require('./routes/web')
const mongoose = require('mongoose')
const { queueInit } = require('./queue')


const api = express()
const web = express()

log4js.configure('./config/log4js.json')

const logger = log4js.getLogger('system')

apiRouteStuffing(api)
webRouteStuffing(web)

// 运行api server
if (!fs.existsSync('./config/global.json')) {
    logger.error('cannot find config/global.json')
    process.exit(1)
}
// 运行web server
logger.info('GMOps web server listening at \'http://0.0.0.0:' + require('../../config/global.json').web_port + '\'')
web.listen(8080)


// 运行api server
if (!fs.existsSync('./config/mongo.json')) {
    logger.error('cannot find config/mongo.json')
    process.exit(1)
}
const mongoConfig = require('../../config/mongo.json')
mongoose.set('useCreateIndex', true)
mongoose.connect(`mongodb://${mongoConfig.host}:${mongoConfig.port}/${mongoConfig.database}`,
    {
        useNewUrlParser: true,
        useUnifiedTopology: true
    }, function (err) {
    if (err) {
        logger.error(`GMOps API server cannnot connected MongoDB, reason: ${err}`)
        return
    }
    const apiUDS = require('../../config/global.json').unix_socket
    fs.exists(apiUDS, (exists) => {
        if (exists) {
            fs.unlinkSync(apiUDS)
        }
        logger.info('GMOps API server listening at \'%s\'', apiUDS)
        api.listen(apiUDS)
    })
})

// 运行任务队列
queueInit(mongoConfig)
