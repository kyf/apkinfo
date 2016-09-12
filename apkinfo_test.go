package apkinfo

import (
	"log"
	"testing"
)

func TestGetApkInfo(t *testing.T) {
	file := "./baidu.apk"
	apk, err := GetApkInfo(file)
	if err != nil {
		t.Error(err)
	}
	log.Print(apk)
}
