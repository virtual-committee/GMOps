const {
    GMOPS_COMMAND_TYPE_ADD_AUTHORIZED,
    SpecCommandArg
} = require('../type')
const fs = require('fs')

class GMOpsAddAuthorizedSpecCommandArg extends SpecCommandArg {
    constructor (args = [], ctx) {
        super(GMOPS_COMMAND_TYPE_ADD_AUTHORIZED, args, true, ctx)
    }

    valid () {
        if (!this.ctx.environment.HOME) {
            return false
        }
        if (this.args.length !== 5) {
            return false
        }
        return true
    }

    exec () {
        const addr = this.args[0]
        const port = this.args[1]
        const userName = this.args[2]
        const keyId = this.args[3]
        const sshKey = this.args[4]
        const sshKeyStore = this.ctx.environment.HOME + '/.ssh'
        if (!fs.existsSync(sshKeyStore)) {
            fs.mkdirSync(sshKeyStore, 0o700)
        }
        const authorizedKeys = sshKeyStore + '/authorized_keys_proxy'
        fs.open(authorizedKeys, 'a', 0o600, function (err, fd) {
            if (err) {
                return console.error(err)
            }
            fs.writeSync(fd, 'command="GWOPS_HOST='
                         + addr
                         + ' GWOPS_PORT='
                         + port
                         +' GMOps/bin/gmops-proxy \''
                         + userName
                         + '\' \''
                         + keyId
                         + '\'", no-port-forwarding,no-X11-forwarding,no-agent-forwarding,no-pty '
                         + sshKey
                         + '\n')
        })
    }
}

module.exports = {
    GMOpsAddAuthorizedSpecCommandArg
}
