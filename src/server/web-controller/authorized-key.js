const { queue } = require('../queue')
const log4js = require('log4js')

const logger = log4js.getLogger('api')

/**
 *
 * 创建应用authorized_key任务
 * @param {String} id authorized_key ID
 * @param {Response} res
 *
 */
function applyAuthorizedKeyEnqueueAction ({ params: { id } }, res) {
    queue.enqueue('applyAuthorizedKeyTask', { id }, function (err, job) {
        logger.info('enqueue apply authorized_key <' + id + '> task') 
    })
    res.status(201).json({ 'status': 'apply authorized task enqueued' }).end()
}

/**
 *
 * 撤销应用authorized_key任务
 * @param {String} id authorized_key ID
 * @param {Response} res
 *
 */
function cancelAuthorizedKeyEnqueueAction ({ params: { id } }, res) {
    queue.enqueue('cancelAuthorizedKeyTask', { id }, function (err, job) {
        logger.info('enqueue cancel authorized_key <' + id + '> task') 
    })
    res.status(202).json({ 'status': 'cancel authorized task enqueued' }).end()
}

module.exports = {
    applyAuthorizedKeyEnqueueAction,
    cancelAuthorizedKeyEnqueueAction
}
