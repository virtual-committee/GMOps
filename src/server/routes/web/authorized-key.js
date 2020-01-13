const express = require('express')
const router = express.Router
const {
    applyAuthorizedKeyEnqueueAction,
    cancelAuthorizedKeyEnqueueAction
} = require('../../web-controller')

module.exports = {
    queue: [
        router().post('/:id/apply', applyAuthorizedKeyEnqueueAction),
        router().post('/:id/cancel', cancelAuthorizedKeyEnqueueAction)
    ]
}
