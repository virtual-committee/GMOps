const {
    SSH_COMMAND_TYPE_NONE,
    SSH_COMMAND_TYPE_GIT_RECEIVE_PACK,

    SpecCommandArg
} = require('./type')
const {
    GitReceiveSpecCommandArg
} = require('./git-spec')

/**
 *
 * 构造git相关的命令
 * @param {List} ssh-command字符串列表
 * @param {Context} 命令执行上下文
 * @return {SpecCommandArg} ssh-command命令
 *
 */
function buildGitSpecCommandArg (command = [], ctx) {
    if (command.length === 0) {
        return new SpecCommandArg()
    }
    switch (command[0]) {
    case 'receive-pack':
        return new GitReceiveSpecCommandArg(command.slice(1), ctx)
    }
    return new SpecCommandArg()
}

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
    case 'git':
        return buildGitSpecCommandArg(command.slice(1), ctx)
    default:
        return new SpecCommandArg()
    }
}

module.exports = exports = {
    buildSpecCommandArg
}
