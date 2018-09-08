package verify

import (
	"net/http"
	"nstream/api"
)

type Verify struct{}

type jsonInput struct {
	Token string `json:"token"`
}

func (v Verify) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/verify"
}

func (v Verify) Handle(request *http.Request, response http.ResponseWriter) {
	var input jsonInput
	jsonErr := api.ReadInput(request, &input)
	if jsonErr != nil {
		api.WriteJsonError(response)
		return
	}
	_, errs := validateInput(input)
	if len(errs) > 0 {
		api.WriteErrors(response, 400, errs)
		return
	}
	verifyUser(input.Token)
}
