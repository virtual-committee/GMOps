const {
    repoModel,
    userRepoModel
} = require('../schemas')
const mongoose = require('mongoose')
const path = require('path')
const fs = require('fs')
const git = require('nodegit')

const repoBasePath = require('../../../config/global.json').repos_base

class Repo {
    constructor ({ _id = new mongoose.Types.ObjectId(), name, descript, attr, user, group }) {
        this.approved = false
        this._id = _id
        this.name = name
        this.descript = descript
        this.attr = attr
        this.user = user
        this.group = group
    }

    _sync ({ _id, name, descript, attr, user, group }) {
        this._id = _id
        this.name = name
        this.descript = descript
        this.attr = attr
        this.user = user
        this.group = group
    }

    /**
     *
     * 创建仓库
     * @return {Boolean} 是否创建成功
     *
     */
    async create () {
        if (typeof this.user === 'undefined' && typeof this.group === 'undefined') {
            return false
        }
        await repoModel.create({
            _id: this._id,
            name: this.name,
            descript: this.descript,
            attr: this.attr
        })
        if (typeof this.user !== 'undefined') {
            await this.chown(this.user)
        }

        const repoPath = path.resolve(repoBasePath, this._id.toString())
        if (fs.existsSync(repoPath)) {
            return false
        }
        fs.mkdirSync(repoPath, 0o700)
        await git.Repository.init(repoPath, 1)
        return true
    }
    
    /**
     *
     * 将当前仓库添加到指定用户中
     * @param {User} user用户
     * @return {Boolean} 添加是否成功
     *
     */
    async chown (user) {
        if (!!(await user.getRepos()).find(repo => repo.name === this.name)) {
            return false
        }
        this.user = user
        await userRepoModel.deleteOne({ _id: this._id })
        await userRepoModel.create({
            user: this.user._id,
            repo: this._id
        })
        return true
    }
}

module.exports = {
    Repo
}
