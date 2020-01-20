#!/bin/bash

protoc --go_out=./ src/proto/error.proto
protoc --go_out=./ src/proto/created.proto
