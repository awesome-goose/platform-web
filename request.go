package web

import (
	"net/http"
)

type Request struct {
	raw *http.Request
}

func NewRequest(raw *http.Request) *Request {
	return &Request{raw}
}

func (r *Request) Read() ([]byte, error) {
	return nil, nil
}
