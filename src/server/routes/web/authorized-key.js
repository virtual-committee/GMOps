const express = require('express')
const router = express.Router
const {
    applyAuthorizedKeyEnqueueAction
} = require('../../web-controller')

module.exports = {
    queue: [
        router().post('/:id/apply', applyAuthorizedKeyEnqueueAction)
    ]
}
