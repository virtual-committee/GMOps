#!/bin/sh

rm -rf bin/build
scripts/build.sh

podman cp bin/build/master/gmops-shell $1:/opt/GMOps/cp
podman exec -it $1 /bin/ash -c "mv /opt/GMOps/cp/gmops-shell /opt/GMOps/gmops-shell \
    && chown -R git:git /opt/GMOps/gmops-shell \
    && rm -rf /opt/GMOps/cp"
