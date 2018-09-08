package verify

import (
	"errors"
	"nstream/data"
)

func validateInput(input jsonInput) (jsonInput, map[string]error) {
	errs := make(map[string]error)
	_, err := validateToken(input.Token)
	if err != nil {
		errs["token"] = err
		return input, errs
	}
	return input, errs
}

func validateToken(token string) (string, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE verification_token = $1)"
	row := data.DB.QueryRow(query, token)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	if !exists {
		return token, errors.New("Invalid token.")
	}
	return token, nil
}
