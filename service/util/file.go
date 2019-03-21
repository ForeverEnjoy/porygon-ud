package util

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	service "github.com/ForeverEnjoy/porygon-ud/service"
)

func CreateFile(filePath string) (isExist bool, err error) {

	dir := path.Dir(filePath)
	log.Printf("dir %s", dir)
	err = os.MkdirAll(dir, os.ModePerm)
	if service.HasError(err) {
		return
	}

	_, err = os.Stat(filePath)
	// create file if not exists
	if nil != err {
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)
			if !service.HasError(err) {
				defer file.Close()
			}
			return false, err
		}
		return
	}

	isExist = true
	return
}

func WriteFile(path string, data string) (err error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if service.HasError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if service.HasError(err) {
		return
	}
	return
}

func ReadFile(path string) (data string, err error) {
	file, err := os.Open(path)
	if service.HasError(err) {
		return
	}

	dataBytes, err := ioutil.ReadAll(file)
	if service.HasError(err) {
		return
	}

	data = string(dataBytes)
	return
}
