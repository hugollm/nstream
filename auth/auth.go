package auth

import (
	"database/sql"
	"net/http"
	"nstream/data"
)

func Authenticate(request *http.Request) (data.User, error) {
	return tokenAuth(request.Header.Get("Auth-Token"))
}

func tokenAuth(token string) (data.User, error) {
	var user data.User
	query := `SELECT users.id, users.email FROM users
	INNER JOIN sessions ON users.id = sessions.user_id
	WHERE sessions.token = $1 LIMIT 1`
	row := data.DB.QueryRow(query, token)
	err := row.Scan(&user.Id, &user.Email)
	if err == sql.ErrNoRows {
		return user, err
	}
	if err != nil {
		panic(err)
	}
	return user, nil
}
