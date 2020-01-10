const express = require('express')
const { authorize } = require('../../middleware/api-authorize')
const bodyParser = require('body-parser')
const router = express.Router
const { 
	userInfoAction,
	userRegisterAction
} = require('../../api-controller')

module.exports = [
	router().use(bodyParser.json()).post('/add', userRegisterAction),
	router().use(authorize).get('/info', userInfoAction)
]
