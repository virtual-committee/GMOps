SHELL=bash
PWD=$(shell pwd)

VERSION?=$(shell git symbolic-ref -q --short HEAD || git describe --all --tags --extact-match 2>/dev/null | awk -F '-' '{print $$NF}')

GOFILES=$(shell find . -type f -name '*.go')
GOFILES_NOVENDOR=$(shell find . -type f -name '*.go' -not -path '*/vendor/*')

SOURCE_DIR?=$(realpath $(firstword $(subst :, ,${GOPATH})))/src/GMOps/src
BUILD_DIR?=$(realpath $(firstword $(subst :, ,${GOPATH})))/src/GMOps/bin/build/$(VERSION)

TARGET_NAME?=

${TARGET_NAME}: $(GOFILES)
	go build -o $(BUILD_DIR)/$(TARGET_NAME)
	@chmod +x $(BUILD_DIR)/$(TARGET_NAME)
