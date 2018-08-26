package logout

import (
	"testing"
	"net/http/httptest"
	"nstream/data/mock"
	"nstream/api"
)

var endpoint Logout = Logout{}

func TestAccept(t *testing.T) {
	request := httptest.NewRequest("POST", "/logout", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func TestSuccessfulLogout(t *testing.T) {
	defer mock.Clear()
	session := mock.Session()
	request := httptest.NewRequest("POST", "/logout", nil)
	response := httptest.NewRecorder()
	request.Header.Add("Auth-Token", session.Token)
	endpoint.Handle(request, response)
	if response.Code != 200 || mock.Exists("sessions", "token", session.Token) {
		t.Fail()
	}
}

func TestAuthError(t *testing.T ) {
	request := httptest.NewRequest("POST", "/logout", nil)
	response := httptest.NewRecorder()
	request.Header.Add("Auth-Token", "invalid-token")
	endpoint.Handle(request, response)
	out := api.NewAuthErrorOutput()
	if response.Code != 401 || response.Body.String() != out.String() {
		t.Fail()
	}
}
