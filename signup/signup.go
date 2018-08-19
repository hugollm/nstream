package signup

import (
    "errors"
    "encoding/json"
    "net/http"
    "net/mail"
    "nts/common"
    "nts/users"
    "strings"
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
    input, jsonErr := readInput(request)
    if jsonErr != nil {
        common.NewJsonError().WriteToResponse(response)
        return
    }
    inputErr := validateInput(&input)
    if inputErr != nil {
        common.NewValidationError(inputErr.Error()).WriteToResponse(response)
        return
    }
    _, err := users.GetUserByEmail(input.Email)
    if err == nil {
        common.NewValidationError("Email is already taken.").WriteToResponse(response)
        return
    }
    users.AddUser(input.Email, input.Password)
}

func readInput(request *http.Request) (SignupInput, error) {
    var input SignupInput
    err := json.NewDecoder(request.Body).Decode(&input)
    return input, err
}

func validateInput(input *SignupInput) error {
    input.Email = strings.TrimSpace(input.Email)
    if input.Email == "" {
        return errors.New("Email is required.")
    }
    email, emailErr := mail.ParseAddress(input.Email)
    if emailErr != nil {
        return errors.New("Invalid email.")
    }
    input.Email = email.Address
    if input.Password == "" {
        return errors.New("Password is required.")
    }
    if len(input.Password) < 8 {
        return errors.New("Password must be at least 8 characters long.")
    }
    return nil
}
