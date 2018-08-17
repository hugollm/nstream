package common

import (
    "fmt"
    "net/http"
)

type Response struct {
    Response http.ResponseWriter
}

func (r Response) WriteString(s string) {
    fmt.Fprint(r.Response, s)
}
