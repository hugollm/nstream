package logout

import (
	"nstream/data"
)

func deleteSession(token string) {
	query := "DELETE FROM sessions WHERE token = $1"
	_, err := data.DB.Exec(query, token)
	if err != nil {
		panic(err)
	}
}
