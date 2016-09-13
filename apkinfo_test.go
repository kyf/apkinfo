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

func TestGetApkInfo1(t *testing.T) {
	file := "./liurenyou-sem001.apk"
	apk, err := GetApkInfo(file)
	if err != nil {
		t.Error(err)
	}
	log.Print(apk)
}
