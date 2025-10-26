#!/bin/bash

set -e

ROOT_DIR="$(pwd)"
VERSION="$(echo `git -C ${ROOT_DIR} describe --abbrev=0 --tags` | sed 's/^.//')" # "v1.2.3" -> "1.2.3"
BUILD_DIR="${ROOT_DIR}/build"
PACKAGE_NAME="aerium-gui_${VERSION}"
PACKAGE_DIR="${ROOT_DIR}/${PACKAGE_NAME}"


# RPM specific metadata
MAINTAINER="Aerium Developers <info@aerium.network>"
LICENSE="MIT"
DESC="Aerium Desktop Wallet (GUI + CLI components)"
URL=$(echo `git remote show origin | grep "Fetch URL"` | awk '{printf $3}')
CATEGORY="User Interface/Desktops"

# Check the architecture
ARC="$(uname -m)"

mkdir -p ${PACKAGE_DIR}

echo "Building the GUI binaries for CentOS ${ARC} architecture"

cd ${ROOT_DIR}
go build -ldflags "-s -w" -trimpath -tags gtk -o ${BUILD_DIR}/aerium-gui ./cmd/gtk


echo "Checking whether FPM is installed"
fpm -v

echo "Building RPM package"

fpm -s dir -t rpm \
	-n "aerium-gui" \
	-v "${VERSION}" \
	-a "${ARC}" \
	--license "${LICENSE}" \
	--maintainer "${MAINTAINER}" \
	--description "${DESC}" \
	--conflicts aerium-cli \
	--replaces aerium-cli \
	--provides aerium-cli \
	--url "${URL}" \
	--category "${CATEGORY}" \
	--rpm-os linux \
	"${BUILD_DIR}/aerium-gui"=/usr/local/bin/aerium-gui \
	"${ROOT_DIR}/.github/releasers/linux/aerium-gui.desktop"=/usr/share/applications/aerium-gui.desktop \
	"${ROOT_DIR}/cmd/gtk/assets/images/logo.png"=/usr/share/icons/hicolor/256x256/apps/aerium.png

echo "RPM package built successfully for ${ARC} architecture"

