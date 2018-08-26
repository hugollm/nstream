package login

import (
	"nstream/data/mock"
	"testing"
)

func TestGetUserByEmail(t *testing.T) {
	defer mock.Clear()
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
	defer mock.Clear()
	mocked := mock.User()
	mock.Update("users", mocked.Id, "email", "JOHN.DOE@gmail.com")
	user, err := getUser("john.doe@GMAIL.COM")
	if err != nil || user.Email != "JOHN.DOE@gmail.com" {
		t.Fail()
	}
}

func TestAddSessionReturnsStrongToken(t *testing.T) {
	defer mock.Clear()
	user := mock.User()
	token := addSession(user.Id)
	if len(token) < 64 {
		t.Fail()
	}
}

func TestAddedSessionGetsPersisted(t *testing.T) {
	defer mock.Clear()
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
