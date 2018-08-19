package signup

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "nts/errors"
    "nts/users"
    "testing"
)

var endpoint Signup = Signup{}

func TestAccept(t *testing.T) {
    request := httptest.NewRequest("POST", "/signup", nil)
    if !endpoint.Accept(request) {
        t.Fail()
    }
    users.ClearUsers()
}

func TestRejectMethod(t *testing.T) {
    request := httptest.NewRequest("GET", "/signup", nil)
    if endpoint.Accept(request) {
        t.Fail()
    }
    users.ClearUsers()
}

func TestRejectPath(t *testing.T) {
    request := httptest.NewRequest("POST", "/signup/", nil)
    if endpoint.Accept(request) {
        t.Fail()
    }
    users.ClearUsers()
}

func TestHandleWithProperData(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "12345678"}
    assertSignup(t, input, 200, nil)
}

func TestNewUserIsRegistered(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "12345678"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    user, err := users.GetUserByEmail("john.doe@gmail.com")
    if err != nil || user.Email != input.Email {
        t.Fail()
    }
    users.ClearUsers()
}

func TestCannotRegisterTheSameEmailTwice(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "12345678"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    err := errors.ValidationError("Email is already taken.")
    assertSignup(t, input, 400, err)
}

func TestInvalidJsonGetsRejected(t *testing.T) {
    body := bytes.NewBuffer([]byte("invalid-json"))
    request := httptest.NewRequest("POST", "/signup", body)
    response := httptest.NewRecorder()
    endpoint.Handle(request, response)
    expectedBody := errors.InvalidJson().Error()
    if response.Code != 400 || response.Body.String() != expectedBody {
        t.Fail()
    }
    users.ClearUsers()
}

func TestEmailIsRequired(t *testing.T) {
    input := SignupInput{"  \n", "12345678"}
    err := errors.ValidationError("Email is required.")
    assertSignup(t, input, 400, err)
}

func TestEmailMustBeValid(t *testing.T) {
    input := SignupInput{"invalid-email", "12345678"}
    err := errors.ValidationError("Invalid email.")
    assertSignup(t, input, 400, err)
}

func TestPasswordIsRequired(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", ""}
    err := errors.ValidationError("Password is required.")
    assertSignup(t, input, 400, err)
}

func TestPasswordCanBeSpaces(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "        "}
    assertSignup(t, input, 200, nil)
}

func TestPasswordCannotBeTooShort(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "1234567"}
    err := errors.ValidationError("Password must be at least 8 characters long.")
    assertSignup(t, input, 400, err)
}

func assertSignup(t *testing.T, input SignupInput, code int, err error) {
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    if response.Code != code {
        t.Fail()
    }
    if err != nil && response.Body.String() != err.Error() {
        t.Fail()
    }
    users.ClearUsers()
}

func makeRequest(input SignupInput) (*http.Request, *httptest.ResponseRecorder) {
    json, _ := json.Marshal(input)
    body := bytes.NewBuffer(json)
    request := httptest.NewRequest("POST", "/signup", body)
    response := httptest.NewRecorder()
    return request, response
}
