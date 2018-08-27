package signup

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"nstream/api"
	"nstream/data/mock"
	"strings"
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
	email := mock.RandString(50) + "@gmail.com"
	input := SignupInput{email, "12345678"}
	assertSignup(t, input, 200, nil)
}

func TestNewUserIsRegistered(t *testing.T) {
	email := mock.RandString(50) + "@gmail.com"
	input := SignupInput{email, "12345678"}
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	if !userWithEmailExists(email) {
		t.Fail()
	}
}

func TestInvalidJsonGetsRejected(t *testing.T) {
	body := strings.NewReader("invalid-json")
	request := httptest.NewRequest("POST", "/signup", body)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
}

func TestSignupInputIsValidated(t *testing.T) {
	input := SignupInput{"", ""}
	errors := map[string]error{
		"email":    errors.New("Email is required."),
		"password": errors.New("Password is required."),
	}
	assertSignup(t, input, 400, errors)
}

func assertSignup(t *testing.T, input SignupInput, code int, errs map[string]error) {
	request, response := makeRequest(input)
	endpoint.Handle(request, response)
	if response.Code != code {
		t.Fail()
	}
	if len(errs) > 0 {
		expected := httptest.NewRecorder()
		api.WriteErrors(expected, 400, errs)
		if response.Code != 400 || response.Body.String() != expected.Body.String() {
			t.Fail()
		}
	}
}

func makeRequest(input SignupInput) (*http.Request, *httptest.ResponseRecorder) {
	body := strings.NewReader(mock.Json(input))
	request := httptest.NewRequest("POST", "/signup", body)
	response := httptest.NewRecorder()
	return request, response
}
