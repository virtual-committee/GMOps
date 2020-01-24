# GMOps
A system like GitOps

# podman运行脚本

```
mkdir -p /opt/GMOps/bin
cp bin/gmops-proxy /opt/GMOps/bin
chmod git:git -R /opt/GMOps
podman run \
    -eGMOPS_MONGO=<mongo_ip> \
    -eGMOPS_GID="$(id -g git)" \
    -eGMOPS_UID="$(id -u git)" \
    -v/home/git/.ssh/authorized_keys:/home/.ssh/host_authorized_keys \
    -d gmops:latest

```
