const shellParser = require('shell-quote').parse
const { buildSpecCommandArg } = require('./cmd/build')
const { Context } = require('./cmd/context')

/**
 *
 * 获取gmops-shell命令
 * @param {Object} environment gmops-shell接收的环境变量
 * @param {Array} argv gmops-shell调用参数，主要用于记录用户principal
 *
 */
function bootstrap (environment = {}, args = process.argv) {
    const ctx = new Context(environment, args)
    if (!ctx.parse()) {
        process.exit(1)
    }
    // 根据环境变量'SSH_ORIGINAL_COMMAND'构造 specCommandArg对象
    const cmd = buildSpecCommandArg(shellParser(environment.SSH_ORIGINAL_COMMAND || '', environment), ctx)
    if (!cmd.valid()) {
        process.exit(2)
    }
    cmd.exec()
}

module.exports = exports = {
    bootstrap
}
