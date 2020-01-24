# GMOps
A system like GitOps

# podman运行脚本

```
mkdir -p /opt/GMOps/bin
cp bin/gmops-proxy /opt/GMOps/bin
chown git:git -R /opt/GMOps
su git -c "ssh-keygen -t rsa -b 4096 -N \"\" -f /home/git/.ssh/id_rsa"
podman run \
    -eGMOPS_MONGO=<mongo_ip> \
    -eGMOPS_GID="$(id -g git)" \
    -eGMOPS_UID="$(id -u git)" \
    -v/home/git/.ssh/authorized_keys:/home/.ssh/host_authorized_keys \
    -v/home/git/.ssh/id_rsa.pub:/home/git/.ssh/authorized_keys \
    -d gmops:latest

```
