#!/usr/bin/env bash
set -ev

dnf clean all && dnf makecache && dnf install -y git python2 make gcc-c++ npm libcurl-devel gettext openssh-server

ssh-keygen -A
groupadd -r git
useradd -d /home/git -g git git

if [ -f "/run/nologin" ]; then 
    rm /run/nologin
fi

npm i -g cnpm
cnpm i
