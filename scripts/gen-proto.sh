#!/bin/bash

CURRENT_DIR=$(pwd)

rm -rf ./genproto/*

for module in $(find $CURRENT_DIR/post-protos/* -type d); do
    protoc -I $CURRENT_DIR/post-protos/ \
           --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/ \
            $module/*.proto;

done;
