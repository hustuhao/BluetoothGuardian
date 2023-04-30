#!/bin/bash

# 设置常量
APP_NAME="BluetoothGuardian"
APP_DIR="cmd/bg"
APP_ICON="Icon.png"
PKG_DIR="package"
DMG_NAME="${APP_NAME}.dmg"
VOL_NAME="${APP_NAME}"
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../" && pwd)"

# 编译并生成可执行文件
cd "${ROOT_DIR}/${APP_DIR}"
go build -o "${APP_NAME}"
fyne package --exe "${APP_NAME}" --profile "${APP_NAME}.app" --name "${APP_NAME}"
cp -R "${APP_NAME}.app" "${ROOT_DIR}/${PKG_DIR}"
rm -rf "${APP_NAME}.app" "${APP_NAME}"

# 创建用于打包的目录
cd "${ROOT_DIR}"
mkdir -p "${PKG_DIR}"
cp "${APP_DIR}/${APP_ICON}" "${PKG_DIR}/${APP_ICON}"

# 切换到打包目录
cd "${ROOT_DIR}/${PKG_DIR}"

# 创建 dmg 安装包
rm -rf "${DMG_NAME}"
hdiutil create -format UDZO -srcfolder "${APP_NAME}.app/" -volname "${VOL_NAME}" "${DMG_NAME}"
rm -rf "${APP_NAME}.app" "${APP_ICON}"
