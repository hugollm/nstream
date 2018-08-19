package signup

import (
    "errors"
    "encoding/json"
    "net/http"
    "net/mail"
    ntserrors "nts/errors"
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
    if err.NotEmpty() {
        err.WriteToResponse(response)
        return
    }
    users.AddUser(input.Email, input.Password)
}

func validateRequest(request *http.Request) (SignupInput, ntserrors.ApiError) {
    input, jsonErr := validateJson(request)
    if jsonErr != nil {
        return input, ntserrors.InvalidJson()
    }
    vEmail, emailErr := validateEmail(input.Email)
    if emailErr != nil {
        return input, ntserrors.ValidationError(emailErr.Error())
    }
    vPassword, passwordErr := validatePassword(input.Password)
    if passwordErr != nil {
        return input, ntserrors.ValidationError(passwordErr.Error())
    }
    return SignupInput{vEmail, vPassword}, ntserrors.Empty()
}

func validateJson(request *http.Request) (SignupInput, error) {
    var input SignupInput
    jsonErr := json.NewDecoder(request.Body).Decode(&input)
    if jsonErr != nil {
        return input, errors.New("Invalid JSON.")
    }
    return input, nil
}

func validateEmail(email string) (string, error) {
    email = strings.TrimSpace(email)
    if email == "" {
        return email, errors.New("Email is required.")
    }
    parsed, parseErr := mail.ParseAddress(email)
    if parseErr != nil {
        return email, errors.New("Invalid email.")
    }
    email = parsed.Address
    _, getErr := users.GetUserByEmail(email)
    if getErr == nil {
        return email, errors.New("Email is already taken.")
    }
    return email, nil
}

func validatePassword(password string) (string, error) {
    if password == "" {
        return password, errors.New("Password is required.")
    }
    if len(password) < 8 {
        return password, errors.New("Password must be at least 8 characters long.")
    }
    return password, nil
}
