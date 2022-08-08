package helpers

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func ResponseJSON(status string, code int, data interface{}) *Response {
	return &Response{Status: status, Code: code, Data: data}
}
