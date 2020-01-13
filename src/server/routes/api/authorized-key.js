const express = require('express')
const router = express.Router
const {
    applyAuthorizedKeyAction,
    cancelAuthorizedKeyAction
} = require('../../api-controller')

module.exports = [
    router().post('/:id/apply', applyAuthorizedKeyAction),
    router().post('/:id/cancel', cancelAuthorizedKeyAction)
]
