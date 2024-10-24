# Android Framework Detector

This repository contains a Go program designed to detect the mobile development framework used in building Android APK and XAPK files. It identifies frameworks such as Flutter, React Native, Cordova, Ionic, Xamarin, and Native (Java/Kotlin).

## Features

- **Framework Detection**: Determines the framework used to build an Android app by inspecting the APK or XAPK file.
- **Supported Frameworks**: 
  - Flutter
  - React Native
  - Cordova
  - Ionic
  - Xamarin
  - Native (Java/Kotlin)
- **XAPK Support**: Automatically processes `.xapk` files, identifying and using the largest `.apk` file contained within for framework detection.

## Installation

You can install the Android Framework Detector in two ways:

### Option 1: Using `go install`

Ensure you have Go installed, then run the following command to install the application directly from the repository:

```bash
go install github.com/AntoineGagnon/android-framework-detector@latest
```

This will download and install the latest version of the Android Framework Detector. After installation, you can run the tool using the command `android-framework-detector`.

### Option 2: Cloning the Repository

1. **Clone the Repository**:

   First, clone the repository to your local machine:

   ```bash
   git clone https://github.com/AntoineGagnon/android-framework-detector.git
   cd android-framework-detector
   ```

2. **Build the Program**:

   Inside the cloned directory, build the application using Go:

   ```bash
   go build -o android-framework-detector main.go
   ```

   This will create an executable named `android-framework-detector` in your directory.

After following either method, you can use the `android-framework-detector` command to run the framework detection tool.

## Usage

Run the program by specifying an APK or XAPK file as an argument.

### Basic Usage

```bash
android-framework-detector <app_name.apk|app_name.xapk>
```

This will output the name of the framework used to build the app.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
 