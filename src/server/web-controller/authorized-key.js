const { queue } = require('../queue')
const log4js = require('log4js')
const {
    applyAuthorizedKey,
    cancelAuthorizedKey
} = require('../web-business')

const logger = log4js.getLogger('api')

/**
 *
 * 创建应用authorized_key任务
 * @param {String} id authorized_key ID
 * @param {Response} res
 *
 */
function applyAuthorizedKeyEnqueueAction ({ params: { id } }, res) {
    applyAuthorizedKey(id)
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
    cancelAuthorizedKey(id)
    res.status(202).json({ 'status': 'cancel authorized task enqueued' }).end()
}

module.exports = {
    applyAuthorizedKeyEnqueueAction,
    cancelAuthorizedKeyEnqueueAction
}
