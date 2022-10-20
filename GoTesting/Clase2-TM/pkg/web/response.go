package web

import "strconv"

//si el campo data o error son nil o vacios los omitirá
type Response struct{
	Code string `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response{
	if code < 300 {
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	return Response{strconv.FormatInt(int64(code), 10), nil, err}
}