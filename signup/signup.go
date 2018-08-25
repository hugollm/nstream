package signup

import (
	"net/http"
	"nstream/api"
)

type Signup struct{}

func (s Signup) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/signup"
}

func (s Signup) Handle(request *http.Request, response http.ResponseWriter) {
	input, errors := validateInput(request.Body)
	if len(errors) > 0 {
		out := api.NewErrorOutput(400, errors)
		out.WriteToResponse(response)
		return
	}
	addUser(input.Email, input.Password)
}
