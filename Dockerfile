FROM alpine:latest

ADD /scripts/install.sh /install.sh
ADD /scripts/start.sh /start.sh

ENV GMOPS_BI_UNIX_SOCKET /opt/GMOps/gmops.sock

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
        && apk --no-cache add git libssh2 libgit2-dev openssh shadow \
        && ln -s /usr/lib/libgit2.so /usr/lib/libgit2.so.27 \
        && mkdir /lib64 \
        && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 \
        && ssh-keygen -A \
        && addgroup git \
        && adduser -h /home/git -G git -D git \
        && sh /install.sh \
        && rm /install.sh

ADD --chown=git:git /bin/build/master /opt/GMOps

EXPOSE 9022

ENTRYPOINT ["/bin/sh", "/start.sh"]
