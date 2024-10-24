package main

import (
	"archive/zip"
	"fmt"
	"os"
	"strings"
)

// Framework represents the name of a mobile development framework.
type Framework string

// Constants for the supported frameworks.
const (
	Flutter    Framework = "Flutter"
	ReactNative          = "React Native"
	Cordova              = "Cordova"
	Ionic                = "Ionic"
	Xamarin              = "Xamarin"
	Native               = "Native (Java/Kotlin)"
)

// Technology holds framework name and associated directory/file paths.
type Technology struct {
	Framework   Framework
	Identifiers []string
}

// techList contains the definitions for each framework's identifying files/directories.
var techList = []Technology{
	{Framework: Flutter, Identifiers: []string{"libflutter.so"}},
	{Framework: ReactNative, Identifiers: []string{"libreactnativejni.so", "assets/index.android.bundle"}},
	{Framework: Cordova, Identifiers: []string{"assets/www/index.html", "assets/www/cordova.js", "assets/www/cordova_plugins.js"}},
	{Framework: Ionic, Identifiers: []string{"assets/native-bridge.js"}},
	{Framework: Xamarin, Identifiers: []string{"/assemblies/Sikur.Monodroid.dll", "/assemblies/Sikur.dll", "/assemblies/Xamarin.Mobile.dll", "/assemblies/mscorlib.dll", "libmonodroid.so", "libmonosgen-2.0.so"}},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: android-framework-detector <app_name.apk>")
		os.Exit(1)
	}

	appName := os.Args[1]
	framework, err := detectFramework(appName)
	if err != nil {
		fmt.Printf("Error detecting framework: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(framework)
}

// detectFramework checks the contents of the APK to determine the framework used.
func detectFramework(appName string) (Framework, error) {
	zipReader, err := zip.OpenReader(appName)
	if err != nil {
		return "", fmt.Errorf("opening zip file: %w", err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		for _, tech := range techList {
			if matchesFramework(file.Name, tech.Identifiers) {
				return tech.Framework, nil
			}
		}
	}
	return Native, nil
}

// matchesFramework checks if a file name matches any of the identifiers for a framework.
func matchesFramework(fileName string, identifiers []string) bool {
	for _, identifier := range identifiers {
		if strings.Contains(fileName, identifier) {
			return true
		}
	}
	return false
}
