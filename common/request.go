package common

import "net/http"

type Request struct {
    Request *http.Request
}

func (r Request) Method() string {
    return r.Request.Method
}

func (r Request) Path() string {
    return r.Request.URL.Path
}
