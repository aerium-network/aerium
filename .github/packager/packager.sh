#!/bin/bash

# The 'set -e' command causes the script to immediately exit
# if any command returns a non-zero exit status (i.e., an error).
set -e

# Helper function for cross-platform 'sed -i'
replace_in_place() {
  if [ "$OSTYPE" = "darwin"* ]; then  # FIX: Changed [[...]] to [...] for sh compatibility
    # macOS requires an empty string argument to -i
    sed -i '' "$1" "$2"
  else
    # Linux (GNU sed) does not require or accept the empty string
    sed -i "$1" "$2"
  fi
}

# --- FIX for "Run from Anywhere" ---
# 1. Determine the ABSOLUTE path to the script itself, regardless of the current working directory.
SCRIPT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# 2. Calculate the repository root (e.g., /mnt/c/Portfolio/aerium)
# The script is expected to be in /path/to/repo/.github/packager, so we go up two levels.
ROOT_DIR=$(dirname $(dirname "$SCRIPT_PATH"))
# --- END FIX ---


PACKAGE_DIR="${ROOT_DIR}/packages"
PROTO_GEN_DIR="${ROOT_DIR}/www/grpc/gen"

if [ -z "$VERSION" ]; then # FIX: Changed [[...]] to [...] for sh compatibility
  echo "❌ Error: Version tag not found."
  exit 1
fi

# Remove 'v' prefix from version if present
VERSION=${VERSION#v}

echo "Packing Version:" ${VERSION}

# --- Cleanup and Global Directory Setup ---
rm -rf ${PACKAGE_DIR}
mkdir -p ${PACKAGE_DIR}
# FIX: Use separate mkdir calls to guarantee that the sh shell creates all folders
mkdir -p ${PACKAGE_DIR}/js/aerium-grpc
mkdir -p ${PACKAGE_DIR}/js/aerium-jsonrpc
mkdir -p ${PACKAGE_DIR}/python/aerium-grpc
mkdir -p ${PACKAGE_DIR}/python/aerium-jsonrpc
mkdir -p ${PACKAGE_DIR}/rust/aerium-grpc
mkdir -p ${PACKAGE_DIR}/rust/aerium-jsonrpc
mkdir -p ${PACKAGE_DIR}/dart/aerium-grpc

# --- Building aerium-grpc package for Dart ---
echo "== Building aerium-grpc package for Dart"
DART_PKG_ROOT="${PACKAGE_DIR}/dart/aerium-grpc"
DART_PROTO_DEST="${DART_PKG_ROOT}/lib/src" # Correct destination for proto files
DART_PUBSPEC_FILE="${DART_PKG_ROOT}/pubspec.yaml" # Added for encoding fix below

# 1. Copy package files (pubspec.yaml, main library file, etc.)
# FIX: Use absolute path relative to the new calculated ROOT_DIR
cp -R "${ROOT_DIR}/.github/packager/dart/." "${DART_PKG_ROOT}/"

# NEW FIX 1: Remove potentially problematic ignore files that hide the pubspec and license.
rm -f "${DART_PKG_ROOT}/.gitignore" "${DART_PKG_ROOT}/.pubignore"

# 2. Create the required 'lib/src' directory for generated protobuf files
mkdir -p "${DART_PROTO_DEST}"

# 3. Copy generated protobuf files into lib/src
cp -R "${PROTO_GEN_DIR}/dart/." "${DART_PROTO_DEST}/"

# 4. FIX: Re-adding encoding fix for pubspec.yaml (Original problem solution)
if command -v iconv >/dev/null 2>&1; then
  echo "   (Fixing pubspec.yaml encoding to UTF-8)"

  # Try converting from UTF-16LE to UTF-8, or fallback to UTF-8 -> UTF-8
  iconv -f UTF-16LE -t UTF-8 "$DART_PUBSPEC_FILE" > temp_pubspec.yaml 2>/dev/null || \
  iconv -f UTF-8 -t UTF-8 "$DART_PUBSPEC_FILE" > temp_pubspec.yaml

  if [ -f temp_pubspec.yaml ]; then
      mv temp_pubspec.yaml "$DART_PUBSPEC_FILE"
  else
      echo "   (Encoding fix failed: temp file not created.)"
  fi
else
  echo "   (Warning: iconv not found, skipping encoding fix. If dry-run fails, install iconv.)"
fi

# 5. Copy license and replace version placeholder
cp ${ROOT_DIR}/LICENSE ${DART_PKG_ROOT}
replace_in_place "s/{{ VERSION }}/$VERSION/g" "$DART_PUBSPEC_FILE"

# NEW FIX 2: Create required documentation files (README and CHANGELOG)
echo "# aerium_grpc" > "${DART_PKG_ROOT}/README.md"
echo "" >> "${DART_PKG_ROOT}/README.md"
echo "Client library for the aerium gRPC API. Generated client for Dart/Flutter projects." >> "${DART_PKG_ROOT}/README.md"
echo "" >> "${DART_PKG_ROOT}/README.md"
echo "## $VERSION" > "${DART_PKG_ROOT}/CHANGELOG.md"
echo "* Initial release." >> "${DART_PKG_ROOT}/CHANGELOG.md"


# --- Building aerium-grpc package for JavaScript ---
echo "== Building aerium-grpc package for JavaScript"
# FIX: Use absolute path relative to the new calculated ROOT_DIR
cp -R ${ROOT_DIR}/.github/packager/js/grpc/* ${PACKAGE_DIR}/js/aerium-grpc
cp -R ${PROTO_GEN_DIR}/js/* ${PACKAGE_DIR}/js/aerium-grpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/js/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/js/aerium-grpc/package.json"

# --- Building aerium-jsonrpc package for JavaScript ---
echo "== Building aerium-jsonrpc package for JavaScript"
GENERATOR_DIR="${PACKAGE_DIR}/generator"
# Note: Cloning inside the package directory. Ensure generator dir is not packaged.
git clone https://github.com/aerium-network/generator.git "$GENERATOR_DIR" && cd "$GENERATOR_DIR"

# FIX: Removed the problematic 'npm config set script-shell /bin/sh'
# We manually perform cleanup and patch package.json to avoid the inner 'rm' command failure.
rm -rf build # Manual cleanup (safe in sh shell)

# Patch the 'build' script in package.json to skip the 'build:clean' step (which uses rm)
# Original: "build": "npm run build:clean && tsc && chmod +x build/cli.js"
# New: "build": "tsc && chmod +x build/cli.js"
# Note: Escaping slashes inside the string for sed.
replace_in_place 's/"build": "npm run build:clean \&\& tsc \&\& chmod +x build\/cli.js"/"build": "tsc \&\& chmod +x build\/cli.js"/' package.json

npm install
npm run build

cd "$ROOT_DIR" && $GENERATOR_DIR/build/cli.js generate \
  -t client \
  -l typescript \
  -n aerium-jsonrpc \
  -d "${ROOT_DIR}/www/grpc/gen/open-rpc/aerium-openrpc.json" \
  -o "$GENERATOR_DIR/js"
cd "$GENERATOR_DIR/js/client/typescript"
npm install

# NEW FIX: Remove 'typedoc' from the build script to bypass the Windows/WSL long path error
# The original build script is usually: "tsc && typedoc ..."
replace_in_place 's/ \&\& typedoc src\/index.ts --out docs --excludeExternals --excludePrivate --hideGenerator --readme none//' package.json

# FIX: Replace direct 'tsc' call with 'npm run build' to use the local tsc binary
npm run build
cp $GENERATOR_DIR/js/client/typescript/build/index.d.ts ${PACKAGE_DIR}/js/aerium-jsonrpc
cp $GENERATOR_DIR/js/client/typescript/build/index.js ${PACKAGE_DIR}/js/aerium-jsonrpc
cp $GENERATOR_DIR/js/client/typescript/build/index.js.map ${PACKAGE_DIR}/js/aerium-jsonrpc
# FIX: Use absolute path relative to the new calculated ROOT_DIR
cp -R ${ROOT_DIR}/.github/packager/js/jsonrpc/* ${PACKAGE_DIR}/js/aerium-jsonrpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/js/aerium-jsonrpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/js/aerium-jsonrpc/package.json"


# --- Building aerium-grpc package for Python ---
echo "== Building aerium-grpc package for Python"
# FIX: Use absolute path relative to the new calculated ROOT_DIR
cp -R ${ROOT_DIR}/.github/packager/python/grpc/* ${PACKAGE_DIR}/python/aerium-grpc
cp ${PROTO_GEN_DIR}/python/* ${PACKAGE_DIR}/python/aerium-grpc/aerium_grpc
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/python/aerium-grpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/python/aerium-grpc/setup.py

# --- Building aerium-jsonrpc package for Python ---
echo "== Building aerium-jsonrpc package for Python"
pip install openrpcclientgenerator

# FIX: Add the user's local bin path to PATH so the 'orpc' command is found.
export PATH="$HOME/.local/bin:$PATH"

ORPC_DIR="${PACKAGE_DIR}/orpc"
mkdir -p ${ORPC_DIR}
cp "${ROOT_DIR}/www/grpc/gen/open-rpc/aerium-openrpc.json" ${ORPC_DIR}/openrpc.json
cd ${ORPC_DIR}
# NEW FIX: Use explicit path to the orpc binary to bypass PATH environment variable issues
~/.local/bin/orpc python example.com ./out
# FIX: Use absolute path relative to the new calculated ROOT_DIR
cp -R ${ROOT_DIR}/.github/packager/python/jsonrpc/* ${PACKAGE_DIR}/python/aerium-jsonrpc
cp ${ORPC_DIR}/out/python/aerium-open-rpc-http-client/aerium_open_rpc_http_client/client.py ${PACKAGE_DIR}/python/aerium-jsonrpc/aerium_jsonrpc/client.py
cp ${ORPC_DIR}/out/python/aerium-open-rpc-http-client/aerium_open_rpc_http_client/models.py ${PACKAGE_DIR}/python/aerium-jsonrpc/aerium_jsonrpc/models.py
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/python/aerium-jsonrpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/python/aerium-jsonrpc/setup.py

# --- Building aerium-grpc package for Rust ---
echo "== Building aerium-grpc package for Rust"
# FIX: Use absolute path relative to the new calculated ROOT_DIR
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
# FIX: Use absolute path relative to the new calculated ROOT_DIR
cp -R ${ROOT_DIR}/.github/packager/rust/jsonrpc/* ${PACKAGE_DIR}/rust/aerium-jsonrpc
cp $GENERATOR_DIR/rust/client/rust/src/index.rs ${PACKAGE_DIR}/rust/aerium-jsonrpc/src/aerium.rs
cp ${ROOT_DIR}/LICENSE ${PACKAGE_DIR}/rust/aerium-jsonrpc
replace_in_place "s/{{ VERSION }}/$VERSION/g" "${PACKAGE_DIR}/rust/aerium-jsonrpc/Cargo.toml"
