package auth

import (
	"net/http/httptest"
	"nstream/data/mock"
	"testing"
)

func TestAuthenticateAcceptsRequest(t *testing.T) {
	session := mock.Session()
	request := httptest.NewRequest("POST", "/mocked", nil)
	request.Header.Add("Auth-Token", session.Token)
	user, err := Authenticate(request)
	if err != nil || user.Id != session.UserId {
		t.Fail()
	}
}

func TestAuthenticateWithNoAuthHeader(t *testing.T) {
	request := httptest.NewRequest("POST", "/mocked", nil)
	_, err := Authenticate(request)
	if err == nil {
		t.Fail()
	}
}

func TestValidTokenAuthReturnsUser(t *testing.T) {
	session := mock.Session()
	user, err := tokenAuth(session.Token)
	if err != nil || user.Id != session.UserId || user.Email == "" {
		t.Fail()
	}
}

func TestInvalidTokenAuthReturnsError(t *testing.T) {
	_, err := tokenAuth("invalid-token")
	if err == nil {
		t.Fail()
	}
}
