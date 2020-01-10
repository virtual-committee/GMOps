const { Principal } = require('../models/principal')

function apiAuthorize (req) {
	req.principal = new Principal(req.get('GMOps-Username'))
	if (!req.principal.valid()) {
		return false
	}
    return true
}

module.exports = {
    apiAuthorize
}
