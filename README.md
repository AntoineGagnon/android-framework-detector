# APK Framework Detector

This repository contains a Go program designed to detect the mobile development framework used in building Android APK files. It identifies frameworks such as Flutter, React Native, Cordova, Ionic, Xamarin, and Native (Java/Kotlin).

## Features

- **Framework Detection**: Determines the framework used to build an Android app by inspecting the APK file.
- **Supported Frameworks**: 
  - Flutter
  - React Native
  - Cordova
  - Ionic
  - Xamarin
  - Native (Java/Kotlin)
- **Verbose Output**: Optionally display detailed output on the framework detection process.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your_username/apk-framework-detector.git
   cd apk-framework-detector
   ```

2. **Build the Program**:
   Ensure you have Go installed, then build the application:
   ```bash
   go build -o apkdetector main.go
   ```

## Usage

Run the program by specifying an APK file as an argument. You can also use the `--verbose` or `-v` flag to get a more detailed output.

### Basic Usage

```bash
./apkdetector <app_name.apk>
```

This will output the framework name directly.

### Verbose Mode

```bash
./apkdetector --verbose <app_name.apk>
# or
./apkdetector -v <app_name.apk>
```

This will output a more detailed statement indicating the framework used.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
 