#!/bin/sh

if [ ! -z $GMOPS_GID ]; then
    echo "reset gid"
    groupmod -g $GMOPS_GID git
fi

if [ ! -z $GMOPS_UID ]; then
    echo "reset uid"
    usermod -u $GMOPS_UID git
fi

/usr/sbin/sshd -p 9022 -D &

if [ ! -d "/home/git/.ssh" ]; then
    mkdir -p /home/git/.ssh
fi

if [ ! -f "/home/git/.ssh/authorized_keys" ]; then
    touch /home/git/.ssh/authorized_keys
fi

if [ -f "/run/nologin" ]; then 
    rm /run/nologin
fi

chown -R git:git /home/git/.ssh
chown -R git:git /opt/GMOps

su git -c "chmod 700 /home/git/.ssh && chmod 600 /home/git/.ssh/authorized_keys"

su git -c "GMOPS_MONGO_HOST=${GMOPS_MONGO_HOST:-127.0.0.1} \
    GMOPS_MONGO_PORT=${GMOPS_MONGO_PORT:-27017} \
    GMOPS_MONGO_DATABASE=${GMOPS_MONGO_DATABASE:-gmops} \
    envsubst < config/mongo.template.json > config/mongo.json"

su git -c "bin/gmops-server-start"
