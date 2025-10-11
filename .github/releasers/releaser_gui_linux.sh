#!/bin/bash

set -e

ROOT_DIR="$(pwd)"
VERSION="$(echo `git -C ${ROOT_DIR} describe --abbrev=0 --tags` | sed 's/^.//')" # "v1.2.3" -> "1.2.3"
BUILD_DIR="${ROOT_DIR}/build"
PACKAGE_NAME="aerium-gui_${VERSION}"
PACKAGE_DIR="${ROOT_DIR}/${PACKAGE_NAME}"

# Check the architecture
ARC="$(uname -m)"

if [ "${ARC}" = "x86_64" ]; then
    FILE_NAME="${PACKAGE_NAME}_linux_amd64"
elif [ "${ARC}" = "aarch64" ]; then
    FILE_NAME="${PACKAGE_NAME}_linux_arm64"
else
    echo "Unsupported architecture: ${ARC}"
    exit 1
fi

mkdir ${PACKAGE_DIR}

echo "Building the binaries for Linux ${ARC} architecture"

cd ${ROOT_DIR}
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/aerium-daemon ./cmd/daemon
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/aerium-wallet ./cmd/wallet
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/aerium-shell ./cmd/shell
go build -ldflags "-s -w" -trimpath -tags gtk -o ${BUILD_DIR}/aerium-gui ./cmd/gtk

# Moving binaries to package directory
echo "Moving binaries"
mv ${BUILD_DIR}/aerium-daemon  ${PACKAGE_DIR}/aerium-daemon
mv ${BUILD_DIR}/aerium-wallet  ${PACKAGE_DIR}/aerium-wallet
mv ${BUILD_DIR}/aerium-shell   ${PACKAGE_DIR}/aerium-shell
mv ${BUILD_DIR}/aerium-gui     ${PACKAGE_DIR}/aerium-gui

echo "Creating archive"
tar -czvf ${ROOT_DIR}/${FILE_NAME}.tar.gz -p ${PACKAGE_NAME}

# building AppImage
# https://github.com/linuxdeploy/linuxdeploy-plugin-gtk

cp ${ROOT_DIR}/.github/releasers/linux/*    ${PACKAGE_DIR}
cp ${ROOT_DIR}/.github/releasers/aerium.png ${PACKAGE_DIR}

cd ${PACKAGE_DIR}

wget -c "https://raw.githubusercontent.com/linuxdeploy/linuxdeploy-plugin-gtk/master/linuxdeploy-plugin-gtk.sh"
wget -c "https://github.com/linuxdeploy/linuxdeploy/releases/download/continuous/linuxdeploy-${ARC}.AppImage"

chmod +x linuxdeploy-${ARC}.AppImage linuxdeploy-plugin-gtk.sh

DEPLOY_GTK_VERSION=3 ./linuxdeploy-${ARC}.AppImage \
    --executable ./aerium-gui \
    --appdir AppDir \
    --plugin gtk \
    --output appimage \
    --icon-file aerium.png \
    --desktop-file ./aerium-gui.desktop

mv ./aerium-gui-${ARC}.AppImage ${ROOT_DIR}/${FILE_NAME}.AppImage
