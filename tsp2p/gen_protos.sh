
#!/bin/bash

set -e

# generate compiles the *.pb.go stubs from the *.proto files.
function generate() {
  echo "Generating root gRPC server protos"
  
  PROTOS=$(find ./protos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
  echo " proto paths ${PROTOS}"
  # For each of the sub-servers, we then generate their protos, but a restricted
  # set as they don't yet require REST proxies, or swagger docs.
  for file in $PROTOS; do
    DIRECTORY=$(dirname "${file}")
    echo "Generating protos from ${file}"
  
    # Generate the protos.
    protoc -I. \
      --go_out=plugins=interfacetype+grpc:. \
      "$(find "${file}" -maxdepth 1 -name '*.proto')"
  
    # Generate the REST reverse proxy.
    protoc -I. \
      --grpc-gateway_out=logtostderr=true,grpc_api_configuration=rest-annotations.yaml:. \
      "$(find "${file}" -maxdepth 1 -name '*.proto')"
  
    # Finally, generate the swagger file which describes the REST API in detail.
    protoc -I. \
      --swagger_out=logtostderr=true:. \
      "$(find "${file}" -maxdepth 1 -name '*.proto')"
  done
}

# Compile and format the lnrpc package.
generate

# move proto files to the right places
cp -r token-strike/tsp2p/* ./server/
rm -rf token-strike


#!/usr/bin/env bash

# see pre-requests:
# - https://grpc.io/docs/languages/go/quickstart/
# - gocosmos plugin is automatically installed during scaffolding.

# set -eo pipefail

# proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
# for dir in $proto_dirs; do
#   protoc \
#   -I "proto" \
#   -I "third_party/proto" \
#   --go_out=plugins=interfacetype+grpc,\
# Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
#   $(find "${dir}" -maxdepth 1 -name '*.proto')

#   # command to generate gRPC gateway (*.pb.gw.go in respective modules) files
#   protoc \
#     -I "proto" \
#     -I "third_party/proto" \
#     --grpc-gateway_out=logtostderr=true:. \
#     $(find "${dir}" -maxdepth 1 -name '*.proto')
# done

# move proto files to the right places
# cp -r token-strike/tsp2p/* ./
# rm -rf token-strike