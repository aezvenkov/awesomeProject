package main

import (
	"errors"
	"github.com/Masterminds/semver"
	"io/ioutil"
	"os"
	"strings"
)

func saveApkFile(fileName string, body []byte) error {
	apkDir := ""

	if strings.Contains(fileName, "Regular") {
		apkDir = "apks/regular/"
	} else if strings.Contains(fileName, "Simulator") {
		apkDir = "apks/simulator/"
	}
	if apkDir == "" {
		return errors.New("invalid APK file name")
	}

	if _, err := os.Stat(apkDir); os.IsNotExist(err) {
		err := os.Mkdir(apkDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	err := ioutil.WriteFile(apkDir+fileName, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

func extractVersion(fileName string) string {
	parts := strings.Split(fileName, "-")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

func compareVersions(version1, version2 string) int {
	v1, err := semver.NewVersion(version1)
	if err != nil {
		return -1
	}
	v2, err := semver.NewVersion(version2)
	if err != nil {
		return 1
	}
	return v1.Compare(v2)
}
