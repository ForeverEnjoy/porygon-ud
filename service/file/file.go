package file

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	service "github.com/ForeverEnjoy/porygon-ud/service"
	service_util "github.com/ForeverEnjoy/porygon-ud/service/util"
)

var basePath = ""

func Init() {
	argBasePath := flag.String("base-path", "", "server url")
	if nil != argBasePath && "" != *argBasePath {
		basePath = *argBasePath
		log.Printf("[File.Info] basePath %s", basePath)
		return
	}

	dir, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	basePath = path.Join(dir, "__porygon_files")
	log.Printf("[File.Info] basePath %s", basePath)
}

type EntityJson struct {
	Id   *string `json:"id"`
	Path *string `json:"path"`
	Data *string `json:"data"`
}

func GetFileBasePath() string {
	return basePath
}

func Post(w http.ResponseWriter, r *http.Request) {
	bodyData, err := service.ReadRequestBody(w, r)
	if nil != err {
		service.MakeErrorCreateEntityResponse(-1, err, w)
		return
	}

	entityJson := &EntityJson{}
	err = service.UnmarshalEntityJson(bodyData, entityJson, w)
	if nil != err {
		service.MakeErrorCreateEntityResponse(-1, err, w)
		return
	}

	if nil == entityJson.Path || nil == entityJson.Data {
		service.MakeErrorCreateEntityResponse(-1, fmt.Errorf("params missing"), w)
		return
	}

	if strings.Contains(*entityJson.Path, "..") {
		service.MakeErrorCreateEntityResponse(-1, fmt.Errorf("invalid path"), w)
		return
	}

	path := filepath.Join(GetFileBasePath(), *entityJson.Path)
	_, err = service_util.CreateFile(path)
	if nil != err {
		service.MakeErrorCreateEntityResponse(1, err, w)
		return
	}

	err = service_util.WriteFile(path, *entityJson.Data)
	if nil != err {
		service.MakeErrorCreateEntityResponse(2, err, w)
		return
	}

	entityJson.Id = &path
	//entityJson.Data = nil
	service.MakeCreateEntityResponse(entityJson, w)
	return
}
