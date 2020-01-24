#!/bin/sh

if [ -f "/run/nologin" ]; then
    rm /run/nologin
fi
