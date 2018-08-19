package signup

import (
    "encoding/json"
    "net/http"
    "nts/users"
)

type Signup struct {}

type SignupInput struct {
    Email string
    Password string
}

func (s Signup) Accept (request *http.Request) bool {
    return request.Method == "POST" && request.URL.Path == "/signup"
}

func (s Signup) Handle (request *http.Request, response http.ResponseWriter) {
    input, err := readInput(request)
    if err != nil {
        response.WriteHeader(400)
        return
    }
    users.AddUser(input.Email, input.Password)
}

func readInput(request *http.Request) (SignupInput, error) {
    var input SignupInput
    err := json.NewDecoder(request.Body).Decode(&input)
    return input, err
}
