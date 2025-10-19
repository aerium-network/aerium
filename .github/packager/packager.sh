#!/bin/bash

# The 'set -e' command causes the script to immediately exit
# if any command returns a non-zero exit status (i.e., an error).
set -e

replace_in_place() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' "$1" "$2"
  else
    sed -i "$1" "$2"
  fi
}

ROOT_DIR="$(pwd)"
PACKAGE_DIR="${ROOT_DIR}/packages"
PROTO_GEN_DIR="${ROOT_DIR}/www/grpc/gen"

if [[ -z "$VERSION" ]]; then
  echo "❌ Error: Version tag not found."
  exit 1
fi

# Remove 'v' prefix from version if present
VERSION=${VERSION#v}

echo "Packing Version:" ${VERSION}

rm -rf ${PACKAGE_DIR}
mkdir -p ${PACKAGE_DIR}
mkdir -p ${PACKAGE_DIR}/js/{aerium-grpc,aerium-jsonrpc}
mkdir -p ${PACKAGE_DIR}/python/{aerium-grpc,aerium-jsonrpc}
mkdir -p ${PACKAGE_DIR}/rust/{aerium-grpc,aerium-jsonrpc}
mkdir -p ${PACKAGE_DIR}/dart/{aerium-grpc,aerium-jsonrpc}

echo "== Building aerium-grpc package for Dart"
cp -R "${ROOT_DIR}/.github/packager/dart/." "${PACKAGE_DIR}/dart/aerium-grpc/"
cp -R "${PROTO_GEN_DIR}/dart/." "${PACKAGE_DIR}/dart/aerium-grpc/lib/"
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/dart/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/dart/aerium-grpc/pubspec.yaml"

echo "== Building aerium-grpc package for JavaScript"
cp -R ${ROOT_DIR}/.github/packager/js/grpc/* ${PACKAGE_DIR}/js/aerium-grpc
cp -R ${PROTO_GEN_DIR}/js/* ${PACKAGE_DIR}/js/aerium-grpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/js/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/js/aerium-grpc/package.json"

echo "== Building aerium-jsonrpc package for JavaScript"
GENERATOR_DIR="${PACKAGE_DIR}/generator"
git clone https://github.com/aerium-network/generator.git "$GENERATOR_DIR" && cd "$GENERATOR_DIR"
npm install && npm run build
cd "$ROOT_DIR" && $GENERATOR_DIR/build/cli.js generate \
  -t client \
  -l typescript \
  -n aerium-jsonrpc \
  -d "${ROOT_DIR}/www/grpc/gen/open-rpc/aerium-openrpc.json" \
  -o "$GENERATOR_DIR/js"
cd "$GENERATOR_DIR/js/client/typescript"
npm install && tsc
cp $GENERATOR_DIR/js/client/typescript/build/index.d.ts ${PACKAGE_DIR}/js/aerium-jsonrpc
cp $GENERATOR_DIR/js/client/typescript/build/index.js ${PACKAGE_DIR}/js/aerium-jsonrpc
cp $GENERATOR_DIR/js/client/typescript/build/index.js.map ${PACKAGE_DIR}/js/aerium-jsonrpc
cp -R ${ROOT_DIR}/.github/packager/js/jsonrpc/* ${PACKAGE_DIR}/js/aerium-jsonrpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/js/aerium-jsonrpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/js/aerium-jsonrpc/package.json"


echo "== Building aerium-grpc package for Python"
cp -R ${ROOT_DIR}/.github/packager/python/grpc/* ${PACKAGE_DIR}/python/aerium-grpc
cp ${PROTO_GEN_DIR}/python/* ${PACKAGE_DIR}/python/aerium-grpc/aerium_grpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/python/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/python/aerium-grpc/setup.py

echo "== Building aerium-jsonrpc package for Python"
pip install openrpcclientgenerator
ORPC_DIR="${PACKAGE_DIR}/orpc"
mkdir -p ${ORPC_DIR}
cp "${ROOT_DIR}/www/grpc/gen/open-rpc/aerium-openrpc.json" ${ORPC_DIR}/openrpc.json
cd ${ORPC_DIR}
orpc python example.com ./out
cp -R ${ROOT_DIR}/.github/packager/python/jsonrpc/* ${PACKAGE_DIR}/python/aerium-jsonrpc
cp ${ORPC_DIR}/out/python/aerium-open-rpc-http-client/aerium_open_rpc_http_client/client.py ${PACKAGE_DIR}/python/aerium-jsonrpc/aerium_jsonrpc/client.py
cp ${ORPC_DIR}/out/python/aerium-open-rpc-http-client/aerium_open_rpc_http_client/models.py ${PACKAGE_DIR}/python/aerium-jsonrpc/aerium_jsonrpc/models.py
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/python/aerium-jsonrpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/python/aerium-jsonrpc/setup.py

echo "== Building aerium-grpc package for Rust"
cp -R ${ROOT_DIR}/.github/packager/rust/grpc/* ${PACKAGE_DIR}/rust/aerium-grpc
cp -R ${PROTO_GEN_DIR}/rust/* ${PACKAGE_DIR}/rust/aerium-grpc/src
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/rust/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/rust/aerium-grpc/Cargo.toml

echo "== Building aerium-jsonrpc package for Rust"
cd "$ROOT_DIR" && $GENERATOR_DIR/build/cli.js generate \
  -t client \
  -l rust \
  -n aerium-jsonrpc \
  -d "${ROOT_DIR}/www/grpc/gen/open-rpc/aerium-openrpc.json" \
  -o "$GENERATOR_DIR/rust"
cp -R ${ROOT_DIR}/.github/packager/rust/jsonrpc/* ${PACKAGE_DIR}/rust/aerium-jsonrpc
cp $GENERATOR_DIR/rust/client/rust/src/index.rs ${PACKAGE_DIR}/rust/aerium-jsonrpc/src/aerium.rs
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/rust/aerium-jsonrpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/rust/aerium-jsonrpc/Cargo.toml"
