#!/bin/bash

set -e

ROOT_DIR="$(pwd)"
VERSION="$(git -C ${ROOT_DIR} describe --abbrev=0 --tags | sed 's/^v//')" # "v1.2.3" -> "1.2.3"
BUILD_DIR="${ROOT_DIR}/build"
PACKAGE_NAME="aerium-cli_${VERSION}"
PACKAGE_DIR="${ROOT_DIR}/${PACKAGE_NAME}"


# RPM specific metadata
MAINTAINER="Aerium Developers <info@aerium.network>"
LICENSE="MIT"
DESC="Aerium Command-Line Tools (daemon, wallet, shell)"
URL=$(git remote get-url origin)
CATEGORY="Applications/Financial"

# Check the architecture
ARC="$(uname -m)"

mkdir -p ${PACKAGE_DIR}

echo "Building the CLI binaries for CentOS ${ARC} architecture"

cd ${ROOT_DIR}
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/aerium-daemon ./cmd/daemon
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/aerium-wallet ./cmd/wallet
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/aerium-shell ./cmd/shell


echo "Checking whether FPM is installed"
fpm -v

echo "Building RPM package"

fpm -s dir -t rpm \
	-n "aerium-cli" \
	-v "${VERSION}" \
	-a "${ARC}" \
	--license "${LICENSE}" \
	--maintainer "${MAINTAINER}" \
	--description "${DESC}" \
	--url "${URL}" \
	--category "${CATEGORY}" \
	--rpm-os linux \
	"${BUILD_DIR}/aerium-daemon"=/usr/local/bin/aerium-daemon \
	"${BUILD_DIR}/aerium-shell"=/usr/local/bin/aerium-shell \
	"${BUILD_DIR}/aerium-wallet"=/usr/local/bin/aerium-wallet

echo "RPM package built successfully for ${ARC} architecture"

