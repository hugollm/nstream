package common

import "net/http"

type Response struct {
    Response http.ResponseWriter
}

func (r Response) WriteString(s string) {
    r.Response.Write([]byte(s))
}
