package status

import (
	"net/http"
	"nstream/api"
)

type Status struct{}

type jsonOutput struct {
	Version string `json:"version"`
}

func (s Status) Accept(request *http.Request) bool {
	return request.Method == "GET" && request.URL.Path == "/"
}

func (s Status) Handle(request *http.Request, response http.ResponseWriter) {
	out := jsonOutput{getVersion()}
	api.WriteOutput(response, 200, out)
}
