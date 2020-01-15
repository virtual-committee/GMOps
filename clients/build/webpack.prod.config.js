const merge = require('webpack-merge')
const path = require('path')
const baseConf = require('./webpack.base.config')
const  { CleanWebpackPlugin } = require('clean-webpack-plugin')
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
module.exports =merge(baseConf, {
    mode: 'production',
    devtool: 'source-map',
    module: {
        rules: []
    },
    plugins: [
      // 新版本的插件只接受一个对象配置作为参数
      new CleanWebpackPlugin({
          root: path.resolve(__dirname, '../dist/')
      }),
      new MiniCssExtractPlugin({
        filename: '[name].css'
      })
    ]
})
