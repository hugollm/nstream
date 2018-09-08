package verify

import (
	"nstream/data"
)

func verifyUser(token string) {
	query := "UPDATE users SET verified = true WHERE verification_token = $1"
	_, err := data.DB.Exec(query, token)
	if err != nil {
		panic(err)
	}
}
