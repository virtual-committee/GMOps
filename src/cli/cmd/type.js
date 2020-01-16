const SSH_COMMAND_TYPE_NONE = 'none'
const SSH_COMMAND_TYPE_GIT_RECEIVE_PACK = 'git-receive-pack'
const SSH_COMMAND_TYPE_GIT_UPLOAD_PACK = 'git-upload-pack'

const GMOPS_COMMAND_TYPE_ADD_AUTHORIZED = 'gmops-add-authorized'
const GMOPS_COMMAND_TYPE_USER_INFO = 'gmops-user-info'

class SpecCommandArg {
    constructor (commandType = SSH_COMMAND_TYPE_NONE, args = [], gmopsCommand = false, ctx) {
        this.commandType = commandType
        this.args = args
        this.gmopsCommand = gmopsCommand
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

    GMOPS_COMMAND_TYPE_USER_INFO,

    SpecCommandArg
}
