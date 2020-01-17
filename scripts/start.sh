#!/bin/sh

/usr/sbin/sshd -p 9022 -D &
su git -c "envsubst < config/mongo.template.json > config/mongo.json"
su git -c "bin/gmops-server-start"
