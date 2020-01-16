const {
    SSH_COMMAND_TYPE_GIT_RECEIVE_PACK,
    SpecCommandArg
} = require('../type')


class GitReceivePackSpecCommandArg extends SpecCommandArg {
    constructor (args = [], ctx) {
        super(SSH_COMMAND_TYPE_GIT_RECEIVE_PACK, args, false, ctx)
    }

    valid () {
        return this.args.length === 1
    }

    exec () {
        require('kexec')('git', [ 'receive-pack' ].concat(this.args))
    }
}

module.exports = {
    GitReceivePackSpecCommandArg
}
