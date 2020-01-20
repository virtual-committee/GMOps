#!/bin/bash

ls src/proto/*.proto | xargs -I {} protoc --go_out=./ {}
