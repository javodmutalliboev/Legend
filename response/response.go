package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   any    `json:"data"`
}

func NewResponse(status string, code int, data any) *Response {
	return &Response{
		Status: status,
		Code:   code,
		Data:   data,
	}
}

func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}
