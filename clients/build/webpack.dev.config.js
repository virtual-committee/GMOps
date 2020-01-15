const merge = require('webpack-merge')
const path = require('path')
const baseConf = require('./webpack.base.config')

module.exports = merge(baseConf, {
    mode: 'development',
    devtool: 'souce-map',
    devServer: {
        contentBase: path.resolve(__dirname, '../dist'),
        hot: true,
        port: 9000
    },
})
