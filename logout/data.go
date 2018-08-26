package logout

import (
	"nstream/api"
)

func deleteSession(token string) {
	query := "DELETE FROM sessions WHERE token = $1"
	_, err := api.DB.Exec(query, token)
	if err != nil {
		panic(err)
	}
}
