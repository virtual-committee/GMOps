#!/bin/bash
set -e
pushd $(pwd) > /dev/null
cd ${GOPATH}/src/GMOps/src
DIRS=$(find * -maxdepth 0 -type d)
for dir in $DIRS; do
    MAKEFILES=$(find $dir -name 'Makefile')
    for makefile in $MAKEFILES; do
        target=$(pwd)/$makefile
        if [ -f $target ]; then
            pushd $(pwd) > /dev/null
            cd $(dirname $makefile)
            make -f Makefile
            if [ $? -ne 0 ]; then
                exit
            fi
            popd > /dev/null
        fi
    done
done
popd > /dev/null
