package login

import (
	"golang.org/x/crypto/bcrypt"
	"nstream/data/mock"
	"testing"
)

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

func TestValidLoginInput(t *testing.T) {
	user := mock.User()
	mock.Update("users", user.Id, "password", hashPassword("12345678"))
	input := LoginInput{user.Email, "12345678"}
	user, errs := validateInput(input)
	if len(errs) > 0 || user.Id == 0 {
		t.Fail()
	}
}

func TestEmailNotFound(t *testing.T) {
	input := LoginInput{"unregistered@gmail.com", "12345678"}
	_, errs := validateInput(input)
	if len(errs) == 0 || errs["email"].Error() != "Email not found." {
		t.Fail()
	}
}

func TestPasswordMismatch(t *testing.T) {
	user := mock.User()
	mock.Update("users", user.Id, "email", user.Email)
	input := LoginInput{user.Email, "wrong-password"}
	_, errs := validateInput(input)
	if len(errs) == 0 || errs["password"].Error() != "Wrong password." {
		t.Fail()
	}
}
