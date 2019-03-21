package uploader

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	service "github.com/ForeverEnjoy/porygon-ud/service"
	service_file "github.com/ForeverEnjoy/porygon-ud/service/file"
	service_util "github.com/ForeverEnjoy/porygon-ud/service/util"
)

type FolderUploader struct {
	serverUrl string
}

func NewFolderUploader(serverUrl string) *FolderUploader {
	return &FolderUploader{
		serverUrl: serverUrl,
	}
}

func (this *FolderUploader) Upload(localPath string) (err error) {
	dirPath := filepath.Dir(localPath)
	filepath.Walk(localPath, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			panic(err)
		}

		if info.IsDir() {
			return nil
		}

		relatePath, err := filepath.Rel(dirPath, path)
		if nil != err {
			panic(err)
		}

		this.upload(path, relatePath)
		return nil
	})
	return
}

func (this *FolderUploader) upload(localPath, serverPath string) (err error) {
	data, err := service_util.ReadFile(localPath)
	if service.HasError(err) {
		return
	}

	entityJson := &service_file.EntityJson{
		Path: &serverPath,
		Data: &data,
	}

	response, err := service.Create(this.serverUrl+"/file", entityJson)
	if 0 != response.Code {
		err = fmt.Errorf("upload error, code [%d], message [%s]", response.Code, response.Message)
		return
	}
	log.Printf("[Upload.Info] file [%s] uploaded", localPath)
	return
}
