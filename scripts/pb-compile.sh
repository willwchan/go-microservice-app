#!/usr/bin/env sh

# Install proto3 from source
#  brew install autoconf automake libtool
#  git clone https://github.com/google/protobuf
#  ./autogen.sh ; ./configure ; make ; make install
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples

REPO_ROOT="${REPO_ROOT:-$(cd "$(dirname "$0")/../.." && pwd)}"
PB_PATH="${REPO_ROOT}/go-microservice-app/api/v1/pb/watermark"
PROTO_FILE=${1:-"watermarksvc.proto"}


echo "Generating pb files for ${PROTO_FILE} service"
protoc -I="${PB_PATH}"  "${PB_PATH}/${PROTO_FILE}" --go-grpc_out=. "${PB_PATH}/${PROTO_FILE}"