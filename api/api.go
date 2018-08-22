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
    found := false
    for _, ep := range api.endpoints {
        if ep.Accept(request) {
            found = true
            ep.Handle(request, response)
            break
        }
    }
    if !found {
        err := errors.New("Endpoint not found.")
        out := NewErrorOutput(404, map[string]error{"not_found": err})
        out.WriteToResponse(response)
    }
}
