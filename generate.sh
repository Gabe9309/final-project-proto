#!/bin/sh
set -e

export PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
export PROTOC_OUT_ROOT_DIR="gen"
export TS_OUT_DIR="${PROTOC_OUT_ROOT_DIR}/ts"
export GO_OUT_DIR="${PROTOC_OUT_ROOT_DIR}/golang"

echo "Cleaning up previous generated files..."
rm -rf ${PROTOC_OUT_ROOT_DIR}

echo "Creating output directories..."
mkdir -p ${TS_OUT_DIR} ${GO_OUT_DIR}

echo "Generating Go protobuf files..."
protoc -I proto \
    --go_out=paths=source_relative:${GO_OUT_DIR} \
    --go-grpc_out=paths=source_relative:${GO_OUT_DIR} \
    proto/user/user.proto \
    proto/weather/weather.proto \
    proto/cache/cache.proto \
    proto/mongo/mongo.proto

echo "Generating Angular (JS/TS) protobuf files..."
npx protoc --ts_out ${TS_OUT_DIR} \
 --proto_path \
    proto proto/user/user.proto \
    proto/weather/weather.proto
# protoc -I proto \
#     --plugin="protoc-gen-ts=$(which protoc-gen-ts)" \
#     --js_out="import_style=commonjs,binary:${TS_OUT_DIR}" \
#     --ts_out="service=grpc-web:${TS_OUT_DIR}" \
#     proto/user/user.proto

echo "Generation complete. Listing generated files:"
find gen -type f
