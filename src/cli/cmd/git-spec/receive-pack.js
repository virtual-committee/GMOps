const {
    SSH_COMMAND_TYPE_GIT_RECEIVE_PACK,
    SpecCommandArg
} = require('../type')


class GitReceiveSpecCommandArg extends SpecCommandArg {
    constructor (args = [], ctx) {
        super(SSH_COMMAND_TYPE_GIT_RECEIVE_PACK, args, ctx)
    }

    valid () {
        if (this.args.length !== 1) {
            return false
        }
        return true
    }

    exec () {
        require('kexec')('git', [ 'receive-pack' ].concat(this.args))
    }
}

module.exports = {
    GitReceiveSpecCommandArg
}
