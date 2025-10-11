#!/bin/bash

set -e

ROOT_DIR="$(pwd)"
VERSION="$(echo `git -C ${ROOT_DIR} describe --abbrev=0 --tags` | sed 's/^.//')"
BUILD_DIR="${ROOT_DIR}/build"
PACKAGE_NAME="aerium-gui_${VERSION}"
PACKAGE_DIR="${ROOT_DIR}/${PACKAGE_NAME}"
FILE_NAME="${PACKAGE_NAME}_windows_amd64"

echo "🚀 Starting Aerium GUI Windows packaging..."

# Create package directory
mkdir -p "${PACKAGE_DIR}/aerium-gui"

# Bundle GTK application using Python bundler
# TODO: After active signPath we change unsigned to signed
python3 "${ROOT_DIR}/.github/releasers/windows/gtk-win-bundler.py" \
    "${BUILD_DIR}/unsigned/aerium-gui.exe" \
    "${PACKAGE_DIR}/aerium-gui"

# Move other binaries
cp ${BUILD_DIR}/unsigned/aerium-daemon.exe  ${PACKAGE_DIR}/aerium-daemon.exe
cp ${BUILD_DIR}/unsigned/aerium-wallet.exe  ${PACKAGE_DIR}/aerium-wallet.exe
cp ${BUILD_DIR}/unsigned/aerium-shell.exe   ${PACKAGE_DIR}/aerium-shell.exe
cp ${BUILD_DIR}/unsigned/aerium-gui.exe     ${PACKAGE_DIR}/aerium-gui/aerium-gui.exe

# Create archive
7z a ${ROOT_DIR}/${FILE_NAME}.zip ${PACKAGE_DIR}

# Create installer
cat << EOF > ${ROOT_DIR}/inno.iss
[Setup]
AppId=Aerium
AppName=Aerium
AppVersion=${VERSION}
AppPublisher=Aerium
AppPublisherURL=https://aerium.network
DefaultDirName={autopf}/Aerium
DefaultGroupName=Aerium
SetupIconFile=.github/releasers/aerium.ico
LicenseFile=LICENSE
Uninstallable=yes
UninstallDisplayIcon={app}\\aerium-gui\\aerium-gui.exe

[Files]
Source:"${PACKAGE_NAME}/*"; DestDir:"{app}"; Flags: recursesubdirs

[Icons]
Name:"{group}\\Aerium"; Filename:"{app}\\aerium-gui\\aerium-gui.exe"
Name:"{commondesktop}\\Aerium"; Filename:"{app}\\aerium-gui\\aerium-gui.exe"

[Run]
Filename:"{app}\\aerium-gui\\aerium-gui.exe"; Description:"Launch Aerium"; Flags: postinstall nowait
EOF

# Build installer
INNO_PATH="/c/Program Files (x86)/Inno Setup 6"
INNO_DIR=$(cygpath -w -s '${INNO_PATH}')
"${INNO_DIR}/ISCC.exe" "${ROOT_DIR}/inno.iss"
mv "Output/mysetup.exe" "${BUILD_DIR}/unsigned/${FILE_NAME}_installer.exe"

echo "🎉 Build complete! Package: ${BUILD_DIR}/unsigned/${FILE_NAME}_installer.exe"
