package signup

import (
	"encoding/json"
	"net/http"
	"nstream/api"
)

type Signup struct{}

func (s Signup) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/signup"
}

func (s Signup) Handle(request *http.Request, response http.ResponseWriter) {
	input, err := readInput(request)
	if err != nil {
		out := api.NewJsonErrorOutput()
		out.WriteToResponse(response)
		return
	}
	input, errs := validateInput(input)
	if len(errs) > 0 {
		out := api.NewErrorOutput(400, errs)
		out.WriteToResponse(response)
		return
	}
	addUser(input.Email, input.Password)
}

func readInput(request *http.Request) (SignupInput, error) {
	var input SignupInput
	err := json.NewDecoder(request.Body).Decode(&input)
	return input, err
}
