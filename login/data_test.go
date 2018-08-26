package login

import (
	"nstream/data/mock"
	"strings"
	"testing"
)

func TestGetUserByEmail(t *testing.T) {
	mocked := mock.User()
	user, err := getUser(mocked.Email)
	if err != nil || user.Id == 0 || user.Email != mocked.Email || user.Password != mocked.Password {
		t.Fail()
	}
}

func TestGetUserWithUnregisteredEmail(t *testing.T) {
	_, err := getUser("unregistered@gmail.com")
	if err == nil {
		t.Fail()
	}
}

func TestGetUserIsCaseInsensitive(t *testing.T) {
	mocked := mock.User()
	mocked.Email = strings.ToUpper(mocked.Email)
	mock.Update("users", mocked.Id, "email", mocked.Email)
	user, err := getUser(strings.ToLower(mocked.Email))
	if err != nil || user.Email != mocked.Email {
		t.Fail()
	}
}

func TestAddSessionReturnsStrongToken(t *testing.T) {
	user := mock.User()
	token := addSession(user.Id)
	if len(token) < 64 {
		t.Fail()
	}
}

func TestAddedSessionGetsPersisted(t *testing.T) {
	user := mock.User()
	token := addSession(user.Id)
	if !mock.Exists("sessions", "token", token) {
		t.Fail()
	}
}

func TestMakeToken(t *testing.T) {
	token := makeToken()
	if len(token) < 64 {
		t.Fail()
	}
}
