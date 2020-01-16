const express = require('express')
const { authorize } = require('../../middleware/api-authorize')
const bodyParser = require('body-parser')
const router = express.Router
const { 
    getUserInfoAction,
    userRegisterAction,
    createUserAuthorizedKeyAction,
    getUserReposAction,
    createUserRepoAction
} = require('../../api-controller')

module.exports = [
    router().use(bodyParser.json()).post('/add', userRegisterAction),
    router().use(authorize).get('/info', getUserInfoAction),
    router().use(authorize).use(bodyParser.json()).post('/authorized-key', createUserAuthorizedKeyAction),
    router().use(authorize).use(bodyParser.json()).get('/repos', getUserReposAction),
    router().use(authorize).use(bodyParser.json()).post('/repo/:repoName', createUserRepoAction)
]
