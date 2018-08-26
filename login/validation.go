package login

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"nstream/data"
)

type LoginInput struct {
	Email    string
	Password string
}

func validateInput(input LoginInput) (data.User, map[string]error) {
	var errs = make(map[string]error)
	user, uErr := getUser(input.Email)
	if uErr != nil {
		errs["email"] = errors.New("Email not found.")
		return user, errs
	}
	pErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if pErr != nil {
		errs["password"] = errors.New("Wrong password.")
		return user, errs
	}
	return user, errs
}
