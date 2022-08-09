package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Meta.Code)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
}

func ResponseJSON(message string, code int, status string, data interface{}) *Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	response := Response{
		Meta: meta,
		Data: data,
	}

	return &response
}
