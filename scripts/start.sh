#!/bin/sh

/usr/sbin/sshd -p 9022 -D &
su git -c "GMOPS_MONGO_HOST=${GMOPS_MONGO_HOST:-127.0.0.1} \
    GMOPS_MONGO_PORT=${GMOPS_MONGO_PORT:-27017} \
    GMOPS_MONGO_DATABASE=${GMOPS_MONGO_DATABASE:-gmops} \
    envsubst < config/mongo.template.json > config/mongo.json"
su git -c "bin/gmops-server-start"
