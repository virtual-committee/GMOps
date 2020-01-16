const { Curl } = require('node-libcurl')

/**
 *
 * 对gmops-shell(ssh)中身份信息进行认证
 * @param {String} username 用户ID
 * @param {String} keyID 公钥ID
 *
 */
function authorize (username, keyID) {
    const promise = new Promise(function (resolve, reject) {
        const curl = new Curl()

        curl.setOpt(Curl.option.URL, 'localhost/user/info')
        curl.setOpt(Curl.option.UNIX_SOCKET_PATH, '/var/run/gmops.sock')
        curl.setOpt(Curl.option.HTTPHEADER, ['GMOps-Username: ' + username])
        curl.on('end', function (statusCode, data, headers) {
            let result = false
            do {
                if (statusCode !== 200) {
                    break
                }
                data = JSON.parse(data)
                if (!data.available) {
                    break
                }
                let authorizedKey = data.authorized_keys.find(({ _id }) => _id === keyID)
                if (typeof authorizedKey === 'undefined') {
                    break
                }
                if (!authorizedKey.writed) {
                    break
                }
                result = true
            } while (false)
            this.close()
            resolve(result)
        })

        curl.on('error', curl.close.bind(curl))
        curl.perform()
    })
    return promise
}

module.exports = {
    authorize
}
