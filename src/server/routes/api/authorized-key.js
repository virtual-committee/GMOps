const express = require('express')
const router = express.Router
const {
   applyAuthorizedKeyAction
} = require('../../api-controller')

module.exports = [
    router().post('/:id/apply', applyAuthorizedKeyAction)
]
