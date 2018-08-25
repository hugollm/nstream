package signup

import (
	"encoding/json"
	"errors"
	"net/http"
	"nstream/api"
)

type Signup struct{}

type SignupInput struct {
	Email    string
	Password string
}

type SignupOutput struct {
	Ok bool
}

func (s Signup) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/signup"
}

func (s Signup) Handle(request *http.Request, response http.ResponseWriter) {
	input, errors := validateRequest(request)
	if len(errors) > 0 {
		out := api.NewErrorOutput(400, errors)
		out.WriteToResponse(response)
		return
	}
	addUser(input.Email, input.Password)
}

func validateRequest(request *http.Request) (SignupInput, map[string]error) {
	errors := make(map[string]error)
	input, jsonErr := validateJson(request)
	if jsonErr != nil {
		errors["json"] = jsonErr
		return SignupInput{}, errors
	}
	vEmail, emailErr := validateEmail(input.Email)
	if emailErr != nil {
		errors["email"] = emailErr
	}
	vPassword, passwordErr := validatePassword(input.Password)
	if passwordErr != nil {
		errors["password"] = passwordErr
	}
	return SignupInput{vEmail, vPassword}, errors
}

func validateJson(request *http.Request) (SignupInput, error) {
	var input SignupInput
	jsonErr := json.NewDecoder(request.Body).Decode(&input)
	if jsonErr != nil {
		return input, errors.New("Invalid JSON.")
	}
	return input, nil
}
