const { GMOpsAddAuthorizedSpecCommandArg } = require('./cmd/gmops-spec')
const { Context } = require('./cmd/context')

/**
 *
 * 获取gmops-authorized命令，具体为添加到$HOME/.ssh/authorized_keys中
 * @param {Object} environment gmops-authorized接收的环境变量
 * @param {Array} argv gmops-authorized调用参数
 *
 */
function bootstrap (environment = {}, args = process.argv) {
    const ctx = new Context(environment, args)
    if (!ctx.parse()) {
        process.exit(1)
    }
    const cmd = new GMOpsAddAuthorizedSpecCommandArg(args.slice(1), ctx)
    if (!cmd.valid()) {
        process.exit(2)
    }
    cmd.exec()
}

module.exports = {
    bootstrap
}
