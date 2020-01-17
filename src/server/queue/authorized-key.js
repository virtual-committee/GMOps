const { Curl } = require('node-libcurl')
const log4js = require('log4js')

const logger = log4js.getLogger('queue')

/**
 *
 * 应用authorized_key任务
 * @param {String} id authorized_key id
 * @param {Function} cb 任务回调函数
 *
 */
function applyAuthorizedKeyTask ({ id }, cb) {
    logger.info('processing apply authorized_key, id: ' + id)
    const curl = new Curl()

    curl.setOpt(Curl.option.URL, '127,0.0.1/authorized-key/' + id + '/apply')
    curl.setOpt(Curl.option.HTTPPOST, [])
    curl.setOpt(Curl.option.UNIX_SOCKET_PATH, require('../../../config/global.json').unix_socket)

    curl.on('end', function (statusCode, data, headers) {
        if (statusCode === 201) {
            cb(null, {
                'status': 'success',
                'task': 'apply authorized_key',
                'param': {
                    'id': id
                }
            })
            logger.info('success apply authorized_key, id: ' + id)
        }
        else {
            logger.info('failured apply authorized_key, id: ' + id)
            cb({
                'status_code': statusCode,
                'body': data
            }, {
                'status': 'failured',
                'task': 'apply authorized_key',
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

function cancelAuthorizedKeyTask ({ id }, cb) {
    logger.info('processing cancel authorized_key, id: ' + id)

    curl.setOpt(Curl.option.URL, '127,0.0.1/authorized-key/' + id + '/cancel')
    curl.setOpt(Curl.option.HTTPPOST, [])
    //curl.setOpt(Curl.option.UNIX_SOCKET_PATH, require('../../../config/global.json').unix_socket)

    curl.on('end', function (statusCode, data, headers) {
        if (statusCode === 202) {
            cb(null, {
                'status': 'success',
                'task': 'cancel authorized_key',
                'param': {
                    'id': id
                }
            })
            logger.info('success cancel authorized_key, id: ' + id)
        }
        else {
            logger.info('failured cancel authorized_key, id: ' + id)
            cb({
                'status_code': statusCode,
                'body': data
            }, {
                'status': 'failured',
                'task': 'cancel authorized_key',
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

module.exports = {
    applyAuthorizedKeyTask,
    cancelAuthorizedKeyTask
}
