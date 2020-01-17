const {
    SSH_COMMAND_TYPE_GIT_RECEIVE_PACK,
    SpecCommandArg
} = require('../type')
const kexec = require('kexec')


class GitReceivePackSpecCommandArg extends SpecCommandArg {
    constructor (args = [], ctx) {
        super(SSH_COMMAND_TYPE_GIT_RECEIVE_PACK, args, false, ctx)
    }

    valid () {
        return this.args.length === 1
    }

    exec () {
        (async () => {
            if (!await authorize(this.ctx)) {
                process.exit(1)
            }

            kexec('git', [ 'receive-pack' ].concat(this.args))
        })()
    }
}

module.exports = {
    GitReceivePackSpecCommandArg
}
