#!/bin/sh

rm -rf bin/build
scripts/build.sh

podman cp bin/build/master/gmops-server $1:/opt/GMOps/cp
podman exec -it $1 /bin/ash -c "mv /opt/GMOps/cp/gmops-server /opt/GMOps/gmops-server \
    && chown -R git:git /opt/GMOps/gmops-server \
    && rm -rf /opt/GMOps/cp"

podman restart $1
