const { SpecCommandArg } = require('./type')
const {
    GitReceivePackSpecCommandArg,
    GitUploadPackSpecCommandArg
} = require('./git-spec')
const {
    GMOpsUserInfoSpecCommandArg
} = require('./gmops-spec')

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
            switch (command[1]) {
                case 'upload-pack':
                    return new GitUploadPackSpecCommandArg(command.slice(2), ctx)
                case 'receive-pack':
                    return new GitReceivePackSpecCommandArg(command.slice(2), ctx)
            }
        case 'gmops':
            switch (command[1]) {
                case 'user-info':
                    return new GMOpsUserInfoSpecCommandArg(command.slice(2), ctx)
            }
    }
    return new SpecCommandArg()
}

module.exports = exports = {
    buildSpecCommandArg
}
