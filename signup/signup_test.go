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
    input := SignupInput{Email: "john.doe@gmail.com", Password: "12345678"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    if response.Code != 200 {
        t.Fail()
    }
    users.ClearUsers()
}

func TestNewUserIsRegistered(t *testing.T) {
    email := "john.doe@gmail.com"
    input := SignupInput{Email: email, Password: "12345678"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    user, err := users.GetUserByEmail("john.doe@gmail.com")
    if err != nil || user.Email != email {
        t.Fail()
    }
    users.ClearUsers()
}

func TestCannotRegisterTheSameEmailTwice(t *testing.T) {
    email := "john.doe@gmail.com"
    input := SignupInput{Email: email, Password: "12345678"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    request, response = makeRequest(input)
    endpoint.Handle(request, response)
    expectedBody := errors.ValidationError("Email is already taken.").Error()
    if response.Code != 400 || response.Body.String() != expectedBody {
        t.Fail()
    }
    users.ClearUsers()
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
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    expectedBody := errors.ValidationError("Email is required.").Error()
    if response.Code != 400 || response.Body.String() != expectedBody {
        t.Fail()
    }
    users.ClearUsers()
}

func TestEmailMustBeValid(t *testing.T) {
    input := SignupInput{"invalid-email", "12345678"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    expectedBody := errors.ValidationError("Invalid email.").Error()
    if response.Code != 400 || response.Body.String() != expectedBody {
        t.Fail()
    }
    users.ClearUsers()
}

func TestPasswordIsRequired(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", ""}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    expectedBody := errors.ValidationError("Password is required.").Error()
    if response.Code != 400 || response.Body.String() != expectedBody {
        t.Fail()
    }
    users.ClearUsers()
}

func TestPasswordCanBeSpaces(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "        "}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    if response.Code != 200 {
        t.Fail()
    }
    users.ClearUsers()
}

func TestPasswordCannotBeTooShort(t *testing.T) {
    input := SignupInput{"john.doe@gmail.com", "1234567"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    expectedBody := errors.ValidationError("Password must be at least 8 characters long.").Error()
    if response.Code != 400 || response.Body.String() != expectedBody {
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
