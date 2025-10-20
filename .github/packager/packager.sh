#!/bin/bash

# The 'set -e' command causes the script to immediately exit
# if any command returns a non-zero exit status (i.e., an error).
set -e

# Helper function for cross-platform 'sed -i'
replace_in_place() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS requires an empty string argument to -i
    sed -i '' "$1" "$2"
  else
    # Linux (GNU sed) does not require or accept the empty string
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

# --- Cleanup and Global Directory Setup ---
rm -rf ${PACKAGE_DIR}
mkdir -p ${PACKAGE_DIR}
mkdir -p ${PACKAGE_DIR}/js/{aerium-grpc,aerium-jsonrpc}
mkdir -p ${PACKAGE_DIR}/python/{aerium-grpc,aerium-jsonrpc}
mkdir -p ${PACKAGE_DIR}/rust/{aerium-grpc,aerium-jsonrpc}
mkdir -p ${PACKAGE_DIR}/dart/aerium-grpc

# --- Building aerium-grpc package for Dart ---
echo "== Building aerium-grpc package for Dart"
DART_PKG_ROOT="${PACKAGE_DIR}/dart/aerium-grpc"
DART_PROTO_DEST="${DART_PKG_ROOT}/lib/src" # Correct destination for proto files

# 1. Copy package files (pubspec.yaml, main library file, etc.)
#    The '." and '/' ensures that contents are copied, not the directory itself.
cp -R "${ROOT_DIR}/.github/packager/dart/." "${DART_PKG_ROOT}/"

# 2. Create the required 'lib/src' directory for generated protobuf files
mkdir -p "${DART_PROTO_DEST}"

# 3. Copy generated protobuf files into lib/src
cp -R "${PROTO_GEN_DIR}/dart/." "${DART_PROTO_DEST}/"

# 4. Copy license and replace version placeholder
cp ${ROOT_DIR}/LICENSE ${DART_PKG_ROOT}
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${DART_PKG_ROOT}/pubspec.yaml"

# --- Building aerium-grpc package for JavaScript ---
echo "== Building aerium-grpc package for JavaScript"
cp -R ${ROOT_DIR}/.github/packager/js/grpc/* ${PACKAGE_DIR}/js/aerium-grpc
cp -R ${PROTO_GEN_DIR}/js/* ${PACKAGE_DIR}/js/aerium-grpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/js/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/js/aerium-grpc/package.json"

# --- Building aerium-jsonrpc package for JavaScript ---
echo "== Building aerium-jsonrpc package for JavaScript"
GENERATOR_DIR="${PACKAGE_DIR}/generator"
# Note: Cloning inside the package directory. Ensure generator dir is not packaged.
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


# --- Building aerium-grpc package for Python ---
echo "== Building aerium-grpc package for Python"
cp -R ${ROOT_DIR}/.github/packager/python/grpc/* ${PACKAGE_DIR}/python/aerium-grpc
cp ${PROTO_GEN_DIR}/python/* ${PACKAGE_DIR}/python/aerium-grpc/aerium_grpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/python/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/python/aerium-grpc/setup.py

# --- Building aerium-jsonrpc package for Python ---
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

# --- Building aerium-grpc package for Rust ---
echo "== Building aerium-grpc package for Rust"
cp -R ${ROOT_DIR}/.github/packager/rust/grpc/* ${PACKAGE_DIR}/rust/aerium-grpc
cp -R ${PROTO_GEN_DIR}/rust/* ${PACKAGE_DIR}/rust/aerium-grpc/src
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/rust/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/rust/aerium-grpc/Cargo.toml

# --- Building aerium-jsonrpc package for Rust ---
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
