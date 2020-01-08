const {
    SSH_COMMAND_TYPE_NONE,
    SSH_COMMAND_TYPE_GIT_RECEIVE_PACK,

    SpecCommandArg
} = require('./type')
const {
    GitReceiveSpecCommandArg,
    GitUploadPackSpecCommandArg
} = require('./git-spec')

/**
 *
 * 获取ssh-command的命令
 * @param {List} command ssh-command字符串列表
 * @param {Context} ctx 命令执行上下文
 * @return {SpecArg} ssh-command的命令
 *
 */
function buildSpecCommandArg (command = [], ctx) {
    if (command.length === 0) {
        return new SpecCommandArg()
    }
    switch (command[0]) {
    case 'git-upload-pack':
        return new GitUploadPackSpecCommandArg(command, ctx)
    }
    return new SpecCommandArg()
}

module.exports = exports = {
    buildSpecCommandArg
}
