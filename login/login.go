package login

import (
	"encoding/json"
	"net/http"
	"nstream/api"
)

type Login struct{}

func (login Login) Accept(request *http.Request) bool {
	return request.Method == "POST" && request.URL.Path == "/login"
}

func (login Login) Handle(request *http.Request, response http.ResponseWriter) {
	input, jsonErr := readInput(request)
	if jsonErr != nil {
		out := api.NewJsonErrorOutput()
		out.WriteToResponse(response)
		return
	}
	user, errs := validateInput(input)
	if len(errs) > 0 {
		out := api.NewErrorOutput(400, errs)
		out.WriteToResponse(response)
		return
	}
	token := addSession(user.id)
	out := LoginOutput{token}
	response.Write(out.Json())
}

func readInput(request *http.Request) (LoginInput, error) {
	var input LoginInput
	err := json.NewDecoder(request.Body).Decode(&input)
	return input, err
}

type LoginOutput struct {
	Token string `json:"token"`
}

func (out LoginOutput) Json() []byte {
	jsonOut, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	return jsonOut
}
