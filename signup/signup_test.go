package signup

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "nstream/common"
    "nstream/users"
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
    errors := map[string]error{"email": errors.New("Email is already taken.")}
    assertSignup(t, input, 400, errors)
}

func TestInvalidJsonGetsRejected(t *testing.T) {
    body := bytes.NewBuffer([]byte("invalid-json"))
    request := httptest.NewRequest("POST", "/signup", body)
    response := httptest.NewRecorder()
    endpoint.Handle(request, response)
    errors := map[string]error{"json": errors.New("Invalid JSON.")}
    out := common.NewErrorOutput(400, errors)
    if response.Code != 400 || response.Body.String() != out.String() {
        t.Fail()
    }
    users.ClearUsers()
}

func TestEmailIsRequired(t *testing.T) {
    input := SignupInput{"  \n", "12345678"}
    errors := map[string]error{"email": errors.New("Email is required.")}
    assertSignup(t, input, 400, errors)
}

func TestEmailMustBeValid(t *testing.T) {
    input := SignupInput{"invalid-email", "12345678"}
    errors := map[string]error{"email": errors.New("Invalid email.")}
    assertSignup(t, input, 400, errors)
}

func TestPasswordIsRequired(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", ""}
    errors := map[string]error{"password": errors.New("Password is required.")}
    assertSignup(t, input, 400, errors)
}

func TestPasswordCanBeSpaces(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "        "}
    assertSignup(t, input, 200, nil)
}

func TestPasswordCannotBeTooShort(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "1234567"}
    errors := map[string]error{"password": errors.New("Password must be at least 8 characters long.")}
    assertSignup(t, input, 400, errors)
}

func assertSignup(t *testing.T, input SignupInput, code int, errors map[string]error) {
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    if response.Code != code {
        t.Fail()
    }
    out := common.NewErrorOutput(400, errors)
    if len(errors) > 0 && response.Body.String() != out.String() {
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
