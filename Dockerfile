FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk --no-cache add git libssh2 libgit2-dev \
    && ln -s /usr/lib/libgit2.so /usr/lib/libgit2.so.27 \
    && mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ADD /bin/build/master /opt/GMOps
