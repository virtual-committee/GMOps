class Context {
    constructor (environment = {}, args = []) {
        this.environment = environment
        this.args = args
    }

    parse () {
        return true
    }
}

module.exports = {
    Context
}
