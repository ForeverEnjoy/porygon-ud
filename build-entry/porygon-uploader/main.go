package main

import (
	"flag"
	"log"
	"path/filepath"

	service_uploader "github.com/ForeverEnjoy/porygon-ud/service/uploader"
)

var filePath = flag.String("file", "", "file need upload")
var serviceUrl = flag.String("server-url", "", "server url")

func main() {
	flag.Parse()

	absFilePath, err := filepath.Abs(*filePath)
	if nil != err {
		panic(err)
	}
	log.Printf("%s", absFilePath)

	err = service_uploader.NewFolderUploader(*serviceUrl).Upload(absFilePath)
	if nil != err {
		panic(err)
	}
}
