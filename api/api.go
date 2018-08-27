package api

import (
	"errors"
	"net/http"
)

type Api struct {
	endpoints []Endpoint
}

type Endpoint interface {
	Accept(request *http.Request) bool
	Handle(request *http.Request, response http.ResponseWriter)
}

func NewApi(endpoints []Endpoint) Api {
	return Api{endpoints}
}

func (api Api) Handle(request *http.Request, response http.ResponseWriter) {
	response.Header().Add("Content-Type", "application/json")
	found := api.FindEndpoint(request, response)
	if !found {
		api.NotFound(response)
	}
}

func (api Api) FindEndpoint(request *http.Request, response http.ResponseWriter) bool {
	for _, ep := range api.endpoints {
		if ep.Accept(request) {
			ep.Handle(request, response)
			return true
		}
	}
	return false
}

func (api Api) NotFound(response http.ResponseWriter) {
	err := errors.New("Endpoint not found.")
	WriteErrors(response, 404, map[string]error{"not_found": err})
}
