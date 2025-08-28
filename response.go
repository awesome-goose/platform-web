package web

import (
	"log"
	"net/http"
)

type Response struct {
	raw http.ResponseWriter
}

func NewResponse(raw http.ResponseWriter) *Response {
	return &Response{raw}
}

func (r *Response) Write(data []byte) error {
	log.Println("Response Write:", string(data))
	return nil
}
