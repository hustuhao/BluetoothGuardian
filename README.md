# BluetoothGuardian [中文文档](README.ZH.MD)

BluetoothGuardian is a macOS application that automatically turns off Bluetooth when the screen is turned off and turns it back on when the screen is turned on.

## Background

Every day when I carry my laptop around, I like to connect my Bluetooth headphones to the laptop. However, the headphones automatically connect to the laptop in my backpack every time. I have to turn on the laptop, turn off Bluetooth, and then pair my headphones with my phone. This process is quite troublesome, so I decided to create this application to solve this problem.

## Requirements

- macOS 10.14 or higher
- blueutil (`brew install blueutil`)

## Usage

1. Download the latest release from the [releases](https://github.com/hustuhao/BluetoothGuardian/releases) page.
2. Double-click the downloaded `.dmg` file to mount it.
3. Drag the `BluetoothGuardian` icon to the `Applications` folder.
4. Open `System Preferences` > `Security & Privacy` > `Privacy` > `Accessibility`.
5. Click the lock icon to make changes and enter your password.
6. Click the `+` icon and select `BluetoothGuardian` from the Applications folder.
7. Launch `BluetoothGuardian` from the `Applications` folder.

## Build from Source

To build from source, you need to have Go 1.19 installed.

1. Clone the repository: `https://github.com/hustuhao/BluetoothGuardian.git`
2. Navigate to the repository: `cd BluetoothGuardian`
3. Build the application: `go build -o BluetoothGuardian`
4. Launch the application: `./BluetoothGuardian`

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.