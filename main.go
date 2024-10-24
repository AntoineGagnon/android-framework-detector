package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"log"
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
	verbose := flag.Bool("verbose", false, "display detailed output")
	flag.BoolVar(verbose, "v", false, "display detailed output (shorthand)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Usage: go run main.go [--verbose|-v] <app_name.apk>")
	}

	appName := flag.Arg(0)
	framework, err := detectFramework(appName)
	if err != nil {
		log.Fatalf("Error detecting framework: %v", err)
	}

	if *verbose {
		fmt.Printf("App was written in %s\n", framework)
	} else {
		fmt.Println(framework)
	}
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
