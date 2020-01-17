const {
    GMOPS_COMMAND_TYPE_USER_INFO,

    SpecCommandArg
} = require('../type')
const { Curl } = require('node-libcurl')
const { authorize } = require('../authorize')
const jsome = require('jsome')


class GMOpsUserInfoSpecCommandArg extends SpecCommandArg {
    constructor (args = [], ctx) {
        super(GMOPS_COMMAND_TYPE_USER_INFO, args, true, ctx)
    }

    async valid () {
        return this.ctx.args.length === 3
    }

    exec () {
        (async () => {
            if (!await authorize(this.ctx)) {
                process.exit(1)
            }
            const curl = new Curl()

            curl.setOpt(Curl.option.URL, './user/info')
            curl.setOpt(Curl.option.UNIX_SOCKET_PATH, require('../../../../config/global.json').unix_socket)
            curl.setOpt(Curl.option.HTTPHEADER, ['GMOps-Username: ' + this.ctx.args[1].replace(/^'/, '').replace(/'$/, '')])
            curl.on('end', function (statusCode, data, headers) {
                data = JSON.parse(data)
                jsome(data)
            })

            curl.on('error', curl.close.bind(curl))
            curl.perform()
        })()
    }
}

module.exports = {
    GMOpsUserInfoSpecCommandArg
}
