package verify

import (
	"net/http/httptest"
	"nstream/data"
	"nstream/data/mock"
	"strings"
	"testing"
)

var endpoint = Verify{}

func TestAccept(t *testing.T) {
	request := httptest.NewRequest("POST", "/verify", nil)
	if !endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectMethod(t *testing.T) {
	request := httptest.NewRequest("GET", "/verify", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestRejectPath(t *testing.T) {
	request := httptest.NewRequest("POST", "/verify/", nil)
	if endpoint.Accept(request) {
		t.Fail()
	}
}

func TestHandleVerifiesUser(t *testing.T) {
	user := mock.User()
	token := mock.RandString(64)
	mock.Update("users", user.Id, "verification_token", token)
	body := strings.NewReader(`{"token":"` + token + `"}`)
	request := httptest.NewRequest("POST", "/verify", body)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 200 {
		t.Fail()
	}
	var verified bool
	query := "SELECT verified FROM users WHERE verification_token = $1 LIMIT 1"
	data.DB.QueryRow(query, token).Scan(&verified)
	if !verified {
		t.Fail()
	}
}

func TestHandleJsonError(t *testing.T) {
	request := httptest.NewRequest("POST", "/verify", strings.NewReader("invalid-json"))
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
}

func TestHandleValidationError(t *testing.T) {
	body := strings.NewReader(`{"token":"invalid-token"}`)
	request := httptest.NewRequest("POST", "/verify", body)
	response := httptest.NewRecorder()
	endpoint.Handle(request, response)
	if response.Code != 400 {
		t.Fail()
	}
}
