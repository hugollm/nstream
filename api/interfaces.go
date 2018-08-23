package api

import "net/http"

type Api interface {
	Handle(request *http.Request, response http.ResponseWriter)
}

type Endpoint interface {
	Accept(request *http.Request) bool
	Handle(request *http.Request, response http.ResponseWriter)
}
