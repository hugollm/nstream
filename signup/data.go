package signup

import (
	"database/sql"
	"nstream/data"
)

func userWithEmailExists(email string) bool {
	var id int
	query := "SELECT id FROM users WHERE LOWER(email) = LOWER($1) LIMIT 1"
	err := data.DB.QueryRow(query, email).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	default:
		return id > 0
	}
}

func addUser(email string, pass string) {
	query := "INSERT INTO USERS (email, password) VALUES ($1, $2)"
	_, err := data.DB.Exec(query, email, pass)
	if err != nil {
		panic(err)
	}
}
