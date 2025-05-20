#!/bin/bash

PROTO_DIR="../pingpong"

protoc \
  -I="${PROTO_DIR}" \
  --go_out=../ \
  --go-grpc_out=../ \
  "${PROTO_DIR}/pingpong.proto"
