package login

import (
	"golang.org/x/crypto/bcrypt"
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
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", hashPassword("12345678"))
	input := LoginInput{"john.doe@gmail.com", "12345678"}
	user, errs := validateInput(input)
	if len(errs) > 0 || user.id == 0 {
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
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", hashPassword("12345678"))
	input := LoginInput{"john.doe@gmail.com", "wrong-password"}
	_, errs := validateInput(input)
	if len(errs) == 0 || errs["password"].Error() != "Wrong password." {
		t.Fail()
	}
}
