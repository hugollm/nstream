package api

import (
	"errors"
	"net/http"
)

type NtsApi struct {
	endpoints []Endpoint
}

func NewApi(endpoints []Endpoint) NtsApi {
	return NtsApi{endpoints}
}

func (api NtsApi) Handle(request *http.Request, response http.ResponseWriter) {
	response.Header().Add("Content-Type", "application/json")
	found := api.FindEndpoint(request, response)
	if !found {
		api.NotFound(response)
	}
}

func (api NtsApi) FindEndpoint(request *http.Request, response http.ResponseWriter) bool {
	for _, ep := range api.endpoints {
		if ep.Accept(request) {
			ep.Handle(request, response)
			return true
		}
	}
	return false
}

func (api NtsApi) NotFound(response http.ResponseWriter) {
	err := errors.New("Endpoint not found.")
	out := NewErrorOutput(404, map[string]error{"not_found": err})
	out.WriteToResponse(response)
}
