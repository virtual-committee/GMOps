const { Curl } = require('node-libcurl')
const { queue } = require('../queue')
const log4js = require('log4js')

const logger = log4js.getLogger('api')

/**
 *
 * 创建应用authorized_key任务
 * @param {String} id authorized_key ID
 *
 */
function applyAuthorizedKey (id) {
    queue.enqueue('applyAuthorizedKeyTask', { id }, function (err, job) {
        logger.info('enqueue apply authorized_key <' + id + '> task') 
    })
}

function fastCancelAuthorizedKey (id) {
    const curl = new Curl()

    curl.setOpt(Curl.option.URL, './authorized-key/' + id + '/cancel/fast')
    curl.setOpt(Curl.option.HTTPPOST, [])
    curl.setOpt(Curl.option.UNIX_SOCKET_PATH, '/var/run/gmops.sock')

    curl.on('end', function (statusCode, data, headers) {
        if (statusCode === 202) {
            cb(null, {
                'status': 'success',
                'task': 'fast cancel authorized_key',
                'param': {
                    'id': id
                }
            })
            logger.info('success fast cancel authorized_key, id: ' + id)
        }
        else {
            logger.info('failured fast cancel authorized_key, id: ' + id)
            cb({
                'status_code': statusCode,
                'body': data
            }, {
                'status': 'failured',
                'task': 'fast cancel authorized_key',
                'param': {
                    'id': id
                }
            })
        }
        this.close()
    })

    curl.on('error', curl.close.bind(curl))
    curl.perform()
}

/**
 *
 * 撤销应用authorized_key任务
 * @param {String} id authorized_key ID
 *
 */
function cancelAuthorizedKey (id) {
    fastCancelAuthorizedKey(id)
    queue.enqueue('cancelAuthorizedKeyTask', { id }, function (err, job) {
        logger.info('enqueue cancel authorized_key <' + id + '> task') 
    })
}

module.exports = {
    applyAuthorizedKey,
    cancelAuthorizedKey
}
