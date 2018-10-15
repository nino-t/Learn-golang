package view

import (
	"encoding/json"
	"net/http"
	"sync"
)

type (
	jsonErrorResponse struct {
		Error string `json:"error"`
	}

	jsonDataResponse struct {
		Data interface{} `json:"data"`
	}
)

var (
	jsonErrPool = sync.Pool{
		New: func() interface{} {
			return new(jsonErrorResponse)
		},
	}

	jsonDataPool = sync.Pool{
		New: func() interface{} {
			return new(jsonDataResponse)
		},
	}
)

func (r *jsonErrorResponse) put() {
	jsonErrPool.Put(r)
}

func (r *jsonDataResponse) put() {
	jsonDataPool.Put(r)
}

var mimeJSON = [...]string{"application/json"}

func RenderJSONError(w http.ResponseWriter, errMessage string, statusCode int) {
	h := w.Header()
	h["Content-Type"] = mimeJSON[:]

	response := jsonErrPool.Get().(*jsonErrorResponse)
	response.Error = errMessage

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
	response.put()
}

func RenderJSONData(w http.ResponseWriter, data interface{}, statusCode int) {
	h := w.Header()
	h["Content-Type"] = mimeJSON[:]

	response := jsonDataPool.Get().(*jsonDataResponse)
	response.Data = data

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
	response.put()
}
