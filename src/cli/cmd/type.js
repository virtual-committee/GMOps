const SSH_COMMAND_TYPE_NONE = 'none'
const SSH_COMMAND_TYPE_GIT_RECEIVE_PACK = 'git-receive-pack'
const SSH_COMMAND_TYPE_GIT_UPLOAD_PACK = 'git-upload-pack'


class SpecCommandArg {
    constructor (commandType = SSH_COMMAND_TYPE_NONE, args = [], gwopsCommand = false, ctx) {
        this.commandType = commandType
        this.args = args
        this.gwopsCommand = gwopsCommand
        this.ctx = ctx
    }

    valid () {
        return true
    }

    exec () {
        process.exit(1)
    }
}

module.exports = {
    SSH_COMMAND_TYPE_NONE,
    SSH_COMMAND_TYPE_GIT_RECEIVE_PACK,
    SSH_COMMAND_TYPE_GIT_UPLOAD_PACK,

    SpecCommandArg
}
