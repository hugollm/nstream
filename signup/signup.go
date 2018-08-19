package signup

import (
    "encoding/json"
    "net/http"
    "net/mail"
    "nts/errors"
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
    input, err := validateRequest(request)
    if err.Code != 0 {
        err.WriteToResponse(response)
        return
    }
    users.AddUser(input.Email, input.Password)
}

func validateRequest(request *http.Request) (SignupInput, errors.ApiError) {
    var input SignupInput
    jsonErr := json.NewDecoder(request.Body).Decode(&input)
    if jsonErr != nil {
        return input, errors.InvalidJson()
    }
    input.Email = strings.TrimSpace(input.Email)
    if input.Email == "" {
        return input, errors.ValidationError("Email is required.")
    }
    email, emailErr := mail.ParseAddress(input.Email)
    if emailErr != nil {
        return input, errors.ValidationError("Invalid email.")
    }
    input.Email = email.Address
    if input.Password == "" {
        return input, errors.ValidationError("Password is required.")
    }
    if len(input.Password) < 8 {
        return input, errors.ValidationError("Password must be at least 8 characters long.")
    }
    _, getErr := users.GetUserByEmail(input.Email)
    if getErr == nil {
        return input, errors.ValidationError("Email is already taken.")
    }
    return input, errors.ApiError{}
}
