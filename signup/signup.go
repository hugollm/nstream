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
	var input SignupInput
	err := api.ReadInput(request, &input)
	if err != nil {
		api.WriteJsonError(response)
		return
	}
	input, errs := validateInput(input)
	if len(errs) > 0 {
		api.WriteErrors(response, 400, errs)
		return
	}
	addUser(input.Email, input.Password)
}
