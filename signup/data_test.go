package signup

import (
	"nstream/data/mock"
	"strings"
	"testing"
)

func TestUserWithEmailExistsWithExactEmail(t *testing.T) {
	user := mock.User()
	if userWithEmailExists(user.Email) != true {
		t.Fail()
	}
}

func TestUserWithEmailExistsIsCaseInsensitive(t *testing.T) {
	user := mock.User()
	mock.Update("users", user.Id, "email", strings.ToUpper(user.Email))
	if userWithEmailExists(strings.ToLower(user.Email)) != true {
		t.Fail()
	}
}

func TestAddUserPersistsOnDb(t *testing.T) {
	email := mock.RandString(50)
	addUser(email, "some-hash")
	if !mock.Exists("users", "email", email) {
		t.Fail()
	}
	if !mock.Exists("users", "password", "some-hash") {
		t.Fail()
	}
}
