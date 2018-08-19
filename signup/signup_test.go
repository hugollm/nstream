package signup

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "nts/common"
    "nts/users"
    "testing"
)

var endpoint Signup = Signup{}

func TestAccept(t *testing.T) {
    request := httptest.NewRequest("POST", "/signup", nil)
    if !endpoint.Accept(request) {
        t.Fail()
    }
}

func TestRejectMethod(t *testing.T) {
    request := httptest.NewRequest("GET", "/signup", nil)
    if endpoint.Accept(request) {
        t.Fail()
    }
}

func TestRejectPath(t *testing.T) {
    request := httptest.NewRequest("POST", "/signup/", nil)
    if endpoint.Accept(request) {
        t.Fail()
    }
}

func TestHandleWithProperData(t *testing.T) {
    input := SignupInput{Email: "john.doe@gmail.com", Password: "abc123-"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    if response.Code != 200 {
        t.Fail()
    }
}

func TestNewUserIsRegistered(t *testing.T) {
    email := "john.doe@gmail.com"
    input := SignupInput{Email: email, Password: "abc123-"}
    request, response := makeRequest(input)
    endpoint.Handle(request, response)
    user, err := users.GetUserByEmail("john.doe@gmail.com")
    if err != nil || user.Email != email {
        t.Fail()
    }
}

func TestInvalidJsonGetsRejected(t *testing.T) {
    body := bytes.NewBuffer([]byte("invalid-json"))
    request := httptest.NewRequest("POST", "/signup", body)
    response := httptest.NewRecorder()
    endpoint.Handle(request, response)
    expectedBody := common.NewJsonError().Error()
    if response.Code != 400 || response.Body.String() != expectedBody {
        t.Fail()
    }
}

func makeRequest(input SignupInput) (*http.Request, *httptest.ResponseRecorder) {
    json, _ := json.Marshal(input)
    body := bytes.NewBuffer(json)
    request := httptest.NewRequest("POST", "/signup", body)
    response := httptest.NewRecorder()
    return request, response
}
