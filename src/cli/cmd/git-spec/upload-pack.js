const {
    SSH_COMMAND_TYPE_GIT_UPLOAD_PACK,
    SpecCommandArg
} = require('../type')


class GitUploadPackSpecCommandArg extends SpecCommandArg {
    constructor (args = [], ctx) {
        super(SSH_COMMAND_TYPE_GIT_UPLOAD_PACK, args.slice(1), false, ctx)
    }

    valid () {
        return this.args.length === 1
    }

    exec () {
        require('kexec')('git', [ 'upload-pack' ].concat(this.args))
    }
}

module.exports = {
    GitUploadPackSpecCommandArg
}
