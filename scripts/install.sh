#!/usr/bin/env bash
set -ev

yum update && yum install -y git python2 make gcc-c++ npm libcurl-devel gettext openssh-server

ssh-keygen -A
groupadd -r git
useradd -d /home/git -g git git
chown -R git:git /opt/GMOps

npm i -g cnpm
su git -c "cnpm i"
