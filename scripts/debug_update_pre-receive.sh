#!/bin/sh

rm -rf bin/build
scripts/build.sh

podman cp bin/build/master/pre-receive $1:/opt/GMOps/cp
podman exec -it $1 /bin/ash -c "mv /opt/GMOps/cp/pre-receive /opt/GMOps/pre-receive \
    && chown -R git:git /opt/GMOps/pre-receive \
    && rm -rf /opt/GMOps/cp"
