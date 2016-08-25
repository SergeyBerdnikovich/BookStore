package main

import (
	"encoding/json"
	"net/http"
)

type Endpoints struct{}

func (c Endpoints) RenderJSON(r http.ResponseWriter, status int, v interface{}) {
	var result []byte
	var err error
	result, err = json.Marshal(v)
	if err != nil {
		http.Error(r, err.Error(), 500)
		return
	}
	// json rendered fine, write out the result
	r.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.WriteHeader(status)
	r.Write(result)
}
