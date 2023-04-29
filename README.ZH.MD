# BluetoothGuardian [English](README.MD)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/hustuhao/BluetoothGuardian/blob/main/LICENSE)

BluetoothGuardian 是一个 macOS 应用程序，它可以在屏幕关闭时自动关闭蓝牙，屏幕开启时自动开启蓝牙，方便用户使用蓝牙耳机等设备，避免频繁手动开启关闭蓝牙的操作。

## 背景

每天我背着电脑时，我喜欢用蓝牙耳机连接蓝牙，但是蓝牙耳机每次都是自动连接上了背包里面的电脑，我不得不打开电脑，关闭蓝牙，以让我的蓝牙耳机与手机配对。为了解决这个问题，我开发了 BluetoothGuardian 应用程序，可以自动开关蓝牙，使连接蓝牙更加方便。

## 安装

在终端中执行以下命令来克隆仓库：

```bash
git clone https://github.com/hustuhao/BluetoothGuardian.git
```

在终端中进入仓库目录，执行以下命令来安装依赖：

```bash
brew install blueutil
```

## 使用

在终端中进入仓库目录，执行以下命令来启动应用程序：

```bash
go run main.go
```

应用程序会自动运行在后台，不会有任何界面，可以在屏幕关闭时自动关闭蓝牙，屏幕开启时自动开启蓝牙。

如果需要停止应用程序，可以在终端中按下 `Ctrl+C` 键来终止应用程序。

## 编译

如果需要编译成可执行文件，可以在终端中进入仓库目录，执行以下命令：

```bash
go build -o BluetoothGuardian main.go
```

## 许可证

MIT 许可证，详见 [LICENSE](https://github.com/hustuhao/BluetoothGuardian/blob/main/LICENSE) 文件。

## 贡献

如果您发现了任何 bug 或者有任何建议和改进意见，欢迎在 [GitHub Issues](https://github.com/hustuhao/BluetoothGuardian/issues) 中提交反馈，感谢您的贡献！