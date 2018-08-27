package login

import (
	"net/http"
	"nstream/api"
)

type Login struct{}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string `json:"token"`
}

func (login Login) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/login"
}

func (login Login) Handle(request *http.Request, response http.ResponseWriter) {
	var input LoginInput
	jsonErr := api.ReadInput(request, &input)
	if jsonErr != nil {
		api.WriteJsonError(response)
		return
	}
	user, errs := validateInput(input)
	if len(errs) > 0 {
		api.WriteErrors(response, 400, errs)
		return
	}
	token := addSession(user.Id)
	api.WriteOutput(response, 200, LoginOutput{token})
}
