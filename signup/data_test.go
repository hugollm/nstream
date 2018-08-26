package signup

import (
	"nstream/data/mock"
	"testing"
)

func TestUserWithEmailExistsWithExactEmail(t *testing.T) {
	defer mock.Clear()
	user := mock.User()
	if userWithEmailExists(user.Email) != true {
		t.Fail()
	}
}

func TestUserWithEmailExistsIsCaseInsensitive(t *testing.T) {
	defer mock.Clear()
	user := mock.User()
	mock.Update("users", user.Id, "email", "john.doe@GMAIL.COM")
	if userWithEmailExists("JOHN.DOE@gmail.com") != true {
		t.Fail()
	}
}

func TestAddUserPersistsOnDb(t *testing.T) {
	defer mock.Clear()
	addUser("john.doe@gmail.com", "some-hash")
	if !mock.Exists("users", "email", "john.doe@gmail.com") {
		t.Fail()
	}
	if !mock.Exists("users", "password", "some-hash") {
		t.Fail()
	}
}
