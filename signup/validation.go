package signup

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"strings"
)

func validateEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return email, errors.New("Email is required.")
	}
	parsed, parseErr := mail.ParseAddress(email)
	if parseErr != nil {
		return email, errors.New("Invalid email.")
	}
	email = parsed.Address
	if userWithEmailExists(email) {
		return email, errors.New("Email is already taken.")
	}
	return email, nil
}

func validatePassword(password string) (string, error) {
	if password == "" {
		return password, errors.New("Password is required.")
	}
	if len(password) < 8 {
		return password, errors.New("Password must be at least 8 characters long.")
	}
	return hashPassword(password), nil
}

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}
