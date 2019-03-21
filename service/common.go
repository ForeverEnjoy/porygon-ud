package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (er *errorResponse) setErrorResponse(code int, err error) {
	er.Code = code
	if nil == err {
		er.Message = "success"
	} else {
		er.Message = fmt.Sprintf("%s", err)
	}
}

type BaseResponse struct {
	errorResponse
}

type QueryCollection struct {
	Total    int64       `json:"total"`
	Entities interface{} `json:"entities"`
}

type QueryHighlightCollection struct {
	Total      int64       `json:"total"`
	Entities   interface{} `json:"entities"`
	Highlights interface{} `json:"highlights"`
}

type CreateEntityResponse struct {
	BaseResponse
	Entity interface{} `json:"entity"`
}

func NewCreateEntityResponse(entity interface{}, code int, err error) (out *CreateEntityResponse) {
	out = &CreateEntityResponse{}
	out.Entity = entity
	out.setErrorResponse(code, err)
	return
}

func NewErrorCreateEntityResponseToJson(code int, err error) (out []byte) {
	out, _ = json.Marshal(NewCreateEntityResponse(nil, code, err))
	return
}

type SingleEntityResponse struct {
	BaseResponse
	Entity interface{} `json:"entity"`
}

func NewSingleEntityResponse(entity interface{}, code int, err error) (out *SingleEntityResponse) {
	out = &SingleEntityResponse{}
	out.Entity = entity
	out.setErrorResponse(code, err)
	return
}

func NewErrorSingleEntityResponseToJson(code int, err error) (out []byte) {
	out, _ = json.Marshal(NewSingleEntityResponse(nil, code, err))
	return
}

type PatchEntityResponse struct {
	BaseResponse
	Entity interface{} `json:"entity"`
}

func NewPatchEntityResponse(entity interface{}, code int, err error) (out *PatchEntityResponse) {
	out = &PatchEntityResponse{}
	out.Entity = entity
	out.setErrorResponse(code, err)
	return
}

func NewErrorPatchEntityResponseToJson(code int, err error) (out []byte) {
	out, _ = json.Marshal(NewPatchEntityResponse(nil, code, err))
	return
}

type DeleteEntityResponse struct {
	BaseResponse
	Entity interface{} `json:"entity"`
}

func NewDeleteEntityResponse(entity interface{}, code int, err error) (out *DeleteEntityResponse) {
	out = &DeleteEntityResponse{}
	out.Entity = entity
	out.setErrorResponse(code, err)
	return
}

func NewDeleteEntityResponseToJson(entity interface{}, code int, err error) (out []byte) {
	out, _ = json.Marshal(NewDeleteEntityResponse(entity, code, err))
	return
}

func NewErrorDeleteEntityResponseToJson(code int, err error) (out []byte) {
	out, _ = json.Marshal(NewDeleteEntityResponse(nil, code, err))
	return
}

type QueryEntityResponse struct {
	BaseResponse
	QueryCollection
}

func NewQueryEntityResponse(entities interface{}, total int64, code int, err error) (out *QueryEntityResponse) {
	if nil == entities || reflect.ValueOf(entities).IsNil() {
		entities = make([]interface{}, 0)
	}
	out = &QueryEntityResponse{}
	out.QueryCollection = QueryCollection{
		Total:    total,
		Entities: entities,
	}

	out.setErrorResponse(code, err)
	return
}

type QueryEntityHighlightResponse struct {
	BaseResponse
	QueryHighlightCollection
}

func NewQueryEntityHighlightResponse(entities interface{}, highlights interface{}, total int64, code int, err error) (out *QueryEntityHighlightResponse) {
	if nil == entities || reflect.ValueOf(entities).IsNil() {
		entities = make([]interface{}, 0)
	}
	out = &QueryEntityHighlightResponse{}
	out.QueryHighlightCollection = QueryHighlightCollection{
		Total:      total,
		Entities:   entities,
		Highlights: highlights,
	}

	out.setErrorResponse(code, err)
	return
}

func NewErrorQueryEntityResponseToJson(code int, err error) (out []byte) {
	out, _ = json.Marshal(NewQueryEntityResponse(nil, 0, code, err))
	return
}

func ReadRequestBody(w http.ResponseWriter, r *http.Request) (bodyData []byte, err error) {
	bodyData, err = ioutil.ReadAll(r.Body)
	if nil != err {
		log.Printf("[Service.Error] read request body data error [%s]", err)
		return
	}
	return
}

func UnmarshalEntityJson(bodyData []byte, createJson interface{}, w http.ResponseWriter) (err error) {
	err = json.Unmarshal(bodyData, createJson)
	if nil != err {
		log.Printf("[Service.Error] unmarshal from request body data error [%s]", err)
		return
	}
	return
}

func MakeCreateEntityResponse(responseJson interface{}, w http.ResponseWriter) (err error) {
	createEntityResponse := NewCreateEntityResponse(responseJson, 0, nil)
	err = makeEntityResponse(createEntityResponse, w)
	return
}

func MakeSingleEntityResponse(responseJson interface{}, w http.ResponseWriter) (err error) {
	singleEntityResponse := NewSingleEntityResponse(responseJson, 0, nil)
	err = makeEntityResponse(singleEntityResponse, w)
	return
}

func MakePatchEntityResponse(responseJson interface{}, w http.ResponseWriter) (err error) {
	patchEntityResponse := NewPatchEntityResponse(responseJson, 0, nil)
	err = makeEntityResponse(patchEntityResponse, w)
	return
}

func MakeDeleteEntityResponse(responseJson interface{}, w http.ResponseWriter) (err error) {
	deleteEntityResponse := NewDeleteEntityResponse(responseJson, 0, nil)
	err = makeEntityResponse(deleteEntityResponse, w)
	return
}

func MakeQueryEntityResponse(responseJson interface{}, total int64, w http.ResponseWriter) (err error) {
	queryEntityResponse := NewQueryEntityResponse(responseJson, total, 0, nil)
	err = makeEntityResponse(queryEntityResponse, w)
	return
}

func MakeQueryEntityHighlightResponse(responseEntityJson interface{}, responseHighlightJson interface{}, total int64, w http.ResponseWriter) (err error) {
	queryEntityResponse := NewQueryEntityHighlightResponse(responseEntityJson, responseHighlightJson, total, 0, nil)
	err = makeEntityResponse(queryEntityResponse, w)
	return
}

func MakeErrorCreateEntityResponse(errCode int, errMessage error, w http.ResponseWriter) (err error) {
	createEntityResponse := NewCreateEntityResponse(nil, errCode, errMessage)
	err = makeEntityResponse(createEntityResponse, w)
	return
}

func MakeErrorSingleEntityResponse(errCode int, errMessage error, w http.ResponseWriter) (err error) {
	singleEntityResponse := NewSingleEntityResponse(nil, errCode, errMessage)
	err = makeEntityResponse(singleEntityResponse, w)
	return
}

func MakeErrorPatchEntityResponse(errCode int, errMessage error, w http.ResponseWriter) (err error) {
	patchEntityResponse := NewPatchEntityResponse(nil, errCode, errMessage)
	err = makeEntityResponse(patchEntityResponse, w)
	return
}

func MakeErrorDeleteEntityResponse(errCode int, errMessage error, w http.ResponseWriter) (err error) {
	deleteEntityResponse := NewDeleteEntityResponse(nil, errCode, errMessage)
	err = makeEntityResponse(deleteEntityResponse, w)
	return
}

func MakeErrorQueryEntityResponse(errCode int, errMessage error, w http.ResponseWriter) (err error) {
	queryEntityResponse := NewQueryEntityResponse(nil, 0, errCode, errMessage)
	err = makeEntityResponse(queryEntityResponse, w)
	return
}

func MakeErrorCustomResponse(customResponse interface{}, w http.ResponseWriter) (err error) {
	err = makeEntityResponse(customResponse, w)
	return
}

func makeEntityResponse(entityResponse interface{}, w http.ResponseWriter) (err error) {
	outputData, err := json.Marshal(entityResponse)
	if nil != err {
		log.Printf("[Service.Error] marshal output error [%s]", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(outputData)
	return
}

func Create(url string, entity interface{}) (out *SingleEntityResponse, err error) {
	respBody, err := request(http.MethodPost, url, entity)
	if 0 == len(respBody) {
		return
	}
	out, err = ParseSingleEntityResponse(respBody)
	return
}

func ParseSingleEntityResponse(respBody []byte) (out *SingleEntityResponse, err error) {
	out = &SingleEntityResponse{}
	err = json.Unmarshal(respBody, out)
	if nil != err {
		log.Printf("[Http.Entity.Request.Error] unmarshal single entity response error, %s", err)
		return
	}
	return
}

func request(method, url string, reqBody interface{}) (respBody []byte, err error) {
	var reader *bytes.Buffer
	if nil != reqBody {
		data, err := json.Marshal(reqBody)
		if nil != err {
			log.Printf("[Http.Entity.Request.Error] json marshal error, %s", err)
			return nil, err
		}
		//log.Printf("[Http.Entity.Request.Info] request [%s] body: %s", url, string(data))
		reader = bytes.NewBuffer(data)
	} else {
		reader = bytes.NewBuffer([]byte(""))
	}

	req, err := http.NewRequest(method, url, reader)
	if nil != err {
		log.Printf("[Http.Entity.Request.Error] http NewRequest error, %s", err)
		return
	}

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if nil != err {
		log.Printf("[Http.Entity.Request.Error] Create http do request error, %s", err)
		return
	}

	if 200 != resp.StatusCode {
		err = errors.New("http status error.")
	}

	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Printf("[Http.Entity.Request.Error] read http response body data error, %s", err)
		return
	}
	//log.Printf("[Http.Entity.Request.Info] response [%s] body: %s", url, string(respBody))

	return
}

func HasError(err error) bool {
	if nil != err {
		log.Printf("%s", err)
	}
	return nil != err
}
