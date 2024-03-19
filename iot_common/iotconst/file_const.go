package iotconst

import (
	"os"
	"path/filepath"
	"strings"
)

func GetTemplatesDir() string {
	return strings.Join([]string{GetWorkTempDir(), "templates"}, string(filepath.Separator))
}
func GetBuildRecordDir() string {
	return strings.Join([]string{GetWorkTempDir(), "build_record"}, string(filepath.Separator))
}
func GetKeytoolDir() string {
	return strings.Join([]string{GetWorkTempDir(), "keytool"}, string(filepath.Separator))
}
func GetLaunchScreenDefaultImageDir() string {
	return strings.Join([]string{GetWorkTempDir(), "defaultImage"}, string(filepath.Separator))
}
func GetIosPlistFileDir() string {
	return strings.Join([]string{GetWorkTempDir(), "plist"}, string(filepath.Separator))
}

func GetWorkTempDir() string {
	path, err := os.Getwd()
	if err == nil {
		return strings.Join([]string{path, "temp"}, string(filepath.Separator))
	}
	return os.TempDir()
}
