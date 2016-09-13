package apkinfo

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

type MetaData struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Application struct {
	Meta []MetaData `xml:"meta-data"`
}

type Apkinfo struct {
	VersionCode int         `xml:"versionCode,attr"`
	VersionName string      `xml:"versionName,attr"`
	Package     string      `xml:"package,attr"`
	App         Application `xml:"application"`
}

func unzipManifest(apk string) (data []byte, err error) {
	cf, err := zip.OpenReader(apk)
	if err != nil {
		return
	}
	defer cf.Close()

	var xmldata []byte
	manifest := fmt.Sprintf("tmp_%d", time.Now().UnixNano())
	var reader io.ReadCloser

	for _, f := range cf.File {
		//log.Print(f.Name)
		if strings.EqualFold("AndroidManifest.xml", f.Name) {
			reader, err = f.Open()
			if err != nil {
				return
			}
			xmldata, err = ioutil.ReadAll(reader)
			reader.Close()
			if err != nil {
				return
			}
			err = ioutil.WriteFile(manifest, xmldata, os.ModePerm)
			if err != nil {
				return
			}
			break
		}

	}

	defer func() {
		os.Remove(manifest)
	}()

	cmd := exec.Command("java", "-jar", "AXMLPrinter2.jar", manifest)
	data, err = cmd.Output()
	data = []byte(strings.Replace(string(data), "android:", "", -1))
	return
}

func GetApkInfo(zipfile string) (*Apkinfo, error) {
	xmldata, err := unzipManifest(zipfile)
	if err != nil {
		return nil, err
	}

	var info Apkinfo
	//log.Print(string(xmldata))
	err = xml.Unmarshal(xmldata, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
