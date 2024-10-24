package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Framework string

const (
	Flutter    Framework = "Flutter"
	ReactNative          = "React Native"
	Cordova              = "Cordova"
	Ionic                = "Ionic"
	Xamarin              = "Xamarin"
	Native               = "Native (Java/Kotlin)"
)

type Technology struct {
	Framework   Framework
	Identifiers []string
}

var techList = []Technology{
	{Framework: Flutter, Identifiers: []string{"libflutter.so"}},
	{Framework: ReactNative, Identifiers: []string{"libreactnativejni.so", "assets/index.android.bundle"}},
	{Framework: Cordova, Identifiers: []string{"assets/www/index.html", "assets/www/cordova.js", "assets/www/cordova_plugins.js"}},
	{Framework: Ionic, Identifiers: []string{"assets/native-bridge.js"}},
	{Framework: Xamarin, Identifiers: []string{"/assemblies/Sikur.Monodroid.dll", "/assemblies/Sikur.dll", "/assemblies/Xamarin.Mobile.dll", "/assemblies/mscorlib.dll", "libmonodroid.so", "libmonosgen-2.0.so"}},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: android-framework-detector <app_name.apk|app_name.xapk>")
		os.Exit(1)
	}

	appName := os.Args[1]

	if strings.HasSuffix(appName, ".xapk") {
		largestAPK, err := extractLargestAPK(appName)
		if err != nil {
			fmt.Printf("Error processing XAPK: %v\n", err)
			os.Exit(1)
		}
		appName = largestAPK
	}

	framework, err := detectFramework(appName)
	if err != nil {
		fmt.Printf("Error detecting framework: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(framework)
}

func extractLargestAPK(xapkPath string) (string, error) {
	zipReader, err := zip.OpenReader(xapkPath)
	if err != nil {
		return "", fmt.Errorf("opening xapk file: %w", err)
	}
	defer zipReader.Close()

	var largestAPK string
	var maxSize int64

	for _, file := range zipReader.File {
		if filepath.Ext(file.Name) == ".apk" {
			if file.UncompressedSize64 > uint64(maxSize) {
				maxSize = int64(file.UncompressedSize64)
				largestAPK = file.Name
			}
		}
	}

	if largestAPK == "" {
		return "", fmt.Errorf("no APK file found in XAPK")
	}

	extractedAPKPath := filepath.Join(os.TempDir(), largestAPK)
	if err := extractFile(zipReader, largestAPK, extractedAPKPath); err != nil {
		return "", fmt.Errorf("extracting APK from XAPK: %w", err)
	}

	return extractedAPKPath, nil
}

func extractFile(zipReader *zip.ReadCloser, fileName, destPath string) error {
	for _, file := range zipReader.File {
		if file.Name == fileName {
			srcFile, err := file.Open()
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			return err
		}
	}
	return fmt.Errorf("file %s not found in archive", fileName)
}

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

func matchesFramework(fileName string, identifiers []string) bool {
	for _, identifier := range identifiers {
		if strings.Contains(fileName, identifier) {
			return true
		}
	}
	return false
}
