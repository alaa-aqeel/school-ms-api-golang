package util

import (
	"encoding/json"
	"net/http"
)

type Map map[string]interface{}

type sResponse struct {
	rw         http.ResponseWriter
	statusCode int
}

func Response(w http.ResponseWriter) *sResponse {
	return &sResponse{
		rw:         w,
		statusCode: 200,
	}
}

func (resp *sResponse) StatusCode(code int) *sResponse {
	resp.statusCode = code
	return resp
}

func (resp *sResponse) Json(data any) {
	resp.rw.Header().Set("Content-Type", "Application/json")
	resp.rw.WriteHeader(resp.statusCode)
	json.NewEncoder(resp.rw).Encode(data)
}

func (resp *sResponse) Success(msg string, data any) {
	resp.Json(Map{
		"message": msg,
		"data":    data,
		"status":  "success",
	})
}

func (resp *sResponse) Error(msg string, errors any) {
	resp.Json(Map{
		"message": msg,
		"errors":  errors,
		"status":  "error",
	})
}
