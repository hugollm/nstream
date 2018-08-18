package api

import "net/http"

type Endpoint interface {
    Accept(request *http.Request) bool
    Handle(request *http.Request, response http.ResponseWriter)
}
