const monq = require('monq')
const log4js = require('log4js')

const logger = log4js.getLogger('queue')
const config = require('../../../config/mongo.json')

const jqClient = monq(`mongodb://${config.host}:${config.port}/${config.database}`)

/**
 *
 * 任务队列初始化
 * @param {Object} config
 *
 */
async function queueInit (config) {
    await (async function () {
        const worker = jqClient.worker(['gmops-queue-api'], { collection: config.job_collection })

        worker.register({
            ...require('./authorized-key')
        })
        logger.info('start running task queue')
        worker.start()
    })()
}

const queue = jqClient.queue('gmops-queue-api')

module.exports = {
    queueInit,
    queue
}
