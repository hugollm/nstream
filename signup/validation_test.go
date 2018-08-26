package signup

import (
	"golang.org/x/crypto/bcrypt"
	"nstream/data/mock"
	"testing"
)

func TestValidEmail(t *testing.T) {
	email, err := validateEmail("john.doe@gmail.com")
	if email != "john.doe@gmail.com" || err != nil {
		t.Fail()
	}
}

func TestInvalidEmail(t *testing.T) {
	_, err := validateEmail("invalid-email")
	if err == nil || err.Error() != "Invalid email." {
		t.Fail()
	}
}

func TestEmailIsRequired(t *testing.T) {
	_, err := validateEmail("")
	if err == nil || err.Error() != "Email is required." {
		t.Fail()
	}
}

func TestEmailIsTrimmedBeforeRequiredCheck(t *testing.T) {
	_, err := validateEmail("  \n  ")
	if err == nil || err.Error() != "Email is required." {
		t.Fail()
	}
}

func TestEmailWithNameIsValidButFiltered(t *testing.T) {
	email, err := validateEmail("John Doe <john.doe@gmail.com>")
	if email != "john.doe@gmail.com" || err != nil {
		t.Fail()
	}
}

func TestEmailCasingIsPreserved(t *testing.T) {
	email, err := validateEmail("John.DOE@gmail.com")
	if email != "John.DOE@gmail.com" || err != nil {
		t.Fail()
	}
}

func TestEmailMustNotBeTaken(t *testing.T) {
	defer mock.Clear()
	user := mock.User()
	mock.Update("users", user.Id, "email", "john.doe@gmail.com")
	_, err := validateEmail("John.DOE@gmail.com")
	if err == nil || err.Error() != "Email is already taken." {
		t.Fail()
	}
}

func TestPasswordIsRequired(t *testing.T) {
	_, err := validatePassword("")
	if err == nil || err.Error() != "Password is required." {
		t.Fail()
	}
}

func TestPasswordIsNotTrimmed(t *testing.T) {
	_, err := validatePassword("  3456  ")
	if err != nil {
		t.Fail()
	}
}

func TestPasswordMustNotBeTooShort(t *testing.T) {
	_, err := validatePassword("1234567")
	if err == nil || err.Error() != "Password must be at least 8 characters long." {
		t.Fail()
	}
}

func TestPasswordsAreHashed(t *testing.T) {
	pass, _ := validatePassword("12345678")
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte("12345678"))
	if err != nil {
		t.Fail()
	}
}

func TestHashedPasswordsHaveFixedSizeOf60(t *testing.T) {
	pass, _ := validatePassword("123456789")
	if len(pass) != 60 {
		t.Fail()
	}
}
