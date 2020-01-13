const fs = require('fs')
const os = require('os')
const log4js = require('log4js')

const logger = log4js.getLogger('api')

/**
 *
 * 验证authorize_key
 * @param {AuthorizeKey} authorize_key
 * @return {Object}
 *
 */
async function validAuthorizedKey (authorizedKey) {
    if (!await authorizedKey.load()) {
        return {
            'result': false,
            'status': 404,
            'message': {
                'reason': 'authorized_key dose not exists'
            }
        }
    }
    if (!authorizedKey.approved) {
        return {
            'result': false,
            'status': 404,
            'message': {
                'reason': 'authorized_key dose not exists'
            }
        }
    }
    if (authorizedKey.writed) {
        return {
            'result': false,
            'status': 409,
            'message': {
                'reason': 'already writed'
            }
        }
    }
    return {
        'result': true
    }
}

function createAuthorizedKeysProxyFile () {
    const sshKeyStore = process.env.HOME + '/.ssh'
    if (!fs.existsSync(sshKeyStore)) {
        fs.mkdirSync(sshKeyStore, 0o700)
    }
    return sshKeyStore
}

function getAuthorizedKeysProxyHost () {
    const interfaces = os.networkInterfaces()
    for (let devName in interfaces) {
        const ifce = interfaces[devName].
            find(ifce => ifce.family === 'IPv4'
                && ifce.address !== '127.0.0.1'
                && !ifce.internal)
        if (typeof ifce !== 'undefined') {
            return ifce.address
        }
    }
}

/**
 *
 * 应用authorize_key
 * @param {AuthorizeKey} authorize_key
 * @return {Object}
 *
 */
function applyAuthorizedKey (authorizedKey) {
    if (!fs.existsSync('./config/global.json')) {
        logger.error('cannot find config/global.json')
        return false
    }
    const sshKeyStore = createAuthorizedKeysProxyFile()
    const authorizedKeysFileName = sshKeyStore + '/authorized_keys_proxy'
    
    const host = getAuthorizedKeysProxyHost()
    const port = require('../../../config/global.json').ssh_port
    if (typeof host === 'undefined') {
        logger.error('cannot get host')
        return false
    }

    fs.open(authorizedKeysFileName, 'a', 0o600, function (err, fd) {
        if (err) {
            return
        }
        fs.writeSync(fd, 'command="GWOPS_HOST='
            + host
            + ' GWOPS_PORT='
            + port
            + ' GMOps/bin/gmops-proxy \''
            + authorizedKey.user.username
            + '\' \''
            + authorizedKey._id
            + '\'", no-port-forwarding,no-X11-forwarding,no-agent-forwarding,no-pty '
            + authorizedKey.authorizedKey
            + '\n')
    })
    return true
}

module.exports = {
    validAuthorizedKey,
    applyAuthorizedKey
}
