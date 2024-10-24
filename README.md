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

Run the program by specifying an APK file as an argument.

```bash
./apkdetector <app_name.apk>
```

This will output the framework name directly.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
 