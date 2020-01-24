#!/bin/sh

if [ ! -z $GMOPS_GID ]; then
    echo "gmops reset gid"
    groupmod -g $GMOPS_GID git
fi

if [ ! -z $GMOPS_UID ]; then
    echo "gmops reset uid"
    usermod -u $GMOPS_UID git
fi

if [ ! -d "/home/git/.ssh" ]; then
    mkdir -p /home/git/.ssh
fi

if [ ! -f "/homg/git/.ssh/authorized_keys" ]; then
    touch /home/git/.ssh/authorized_keys
fi

if [ -f "/run/nologin" ]; then
    rm /run/nologin
fi

chown -R git:git /home/git/.ssh
chown -R git:git /opt/GMOps

/usr/sbin/sshd -p 9022 -D &
su git -c "chmod 700 /home/git/.ssh && chmod 600 /home/git/.ssh/authorized_keys"
su git -c "/opt/GMOps/gmops-server --mongo=\"mongodb://${GMOPS_MONGO:-127.0.0.1}\""
