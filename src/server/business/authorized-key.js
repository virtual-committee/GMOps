const fs = require('fs')
const os = require('os')
const log4js = require('log4js')
const lineByLine = require('n-readlines')
const shell = require('shelljs')

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
async function applyAuthorizedKey (authorizedKey) {
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

    const fd = fs.openSync(authorizedKeysFileName, 'a', 0o600)
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
    fs.closeSync(fd)
    authorizedKey.writed = true
    await authorizedKey.save()
    return true
}

function findAuthorizedKeyLineNumber (id) {
    const sshKeyStore = createAuthorizedKeysProxyFile()
    const authorizedKeysFileName = sshKeyStore + '/authorized_keys_proxy'
    const liner = new lineByLine(authorizedKeysFileName)
    let line
    let lineNumber = 0
    while (line = liner.next()) {
        lineNumber += 1
        let times = 0
        let preId = ''
        line = line.toString()
        for (let idx in line) {
            let c = line[idx]
            if (c === '\'') {
                times++
                continue
            }
            if (times === 3) {
                preId += c
            }
            else if (times === 4) {
                break
            }
        }
        if (id === preId) {
            console.log(lineNumber)
            return {
                'result': true,
                'lineNumber': lineNumber,
                'authorizedKeysFileName': authorizedKeysFileName
            }
        }
    }
    return {
        'result': false
    }
}

async function validAuthorizedKeyApplied (authorizedKey) {
    if (!authorizedKey.approved) {
        return {
            'applied': false,
            'message': {
                'reason': 'authorized_key dose not exist'
            }
        }
    }
    const findLineNumberRet = findAuthorizedKeyLineNumber(authorizedKey._id)
    if (!findLineNumberRet.result) {
        return {
            'applied': false,
            'message': {
                'reason': 'authorized_key dose not applied'
            }
        }
    }
    return {
        'applied': true,
        'lineNumber': findLineNumberRet.lineNumber,
        'authorizedKeysFileName': findLineNumberRet.authorizedKeysFileName
    }
}

/**
 *
 * 取消使用authorized_key
 * @param {AuthorizedKey} authorized_key
 * @return {Object}
 *
 */
async function cancelAuthorizedKey (authorizedKey) {
    await authorizedKey.load()
    const validRet = await validAuthorizedKeyApplied(authorizedKey)
    if (!validRet.applied) {
        return {
            'result': false,
            'status': 400,
            'message': {
                ...validRet.message
            }
        }
    }
    const { lineNumber, authorizedKeysFileName } = validRet
    shell.exec('sed -i \'' + lineNumber + 'd\' ' + authorizedKeysFileName)
    authorizedKey.writed = false
    authorizedKey.save()
    return {
        'result': true
    }
}

async function fastCancelAuthorizedKey (authorizedKey) {
    await authorizedKey.load()
    if (!authorizedKey.approved) {
        return {
            'result': false,
            'message': {
                'reason': 'authorized_key dose not exist'
            }
        }
    }
    authorizedKey.writed = false
    authorizedKey.save()
    return {
        'result': true
    }
}

module.exports = {
    validAuthorizedKey,
    applyAuthorizedKey,
    cancelAuthorizedKey,
    fastCancelAuthorizedKey
}
