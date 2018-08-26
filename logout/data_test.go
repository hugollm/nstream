package logout

import (
	"nstream/api"
	"testing"
)

func clearDbUsers() {
	_, err := api.DB.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}

func makeDbUser(email string, pass string) (id int) {
	row := api.DB.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id", email, pass)
	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func makeDbSession(userId int, token string) {
	query := "INSERT INTO sessions (user_id, token) VALUES ($1, $2)"
	_, err := api.DB.Exec(query, userId, token)
	if err != nil {
		panic(err)
	}
}

func dbSessionExists(token string) (exists bool) {
	query := "SELECT EXISTS (SELECT 1 FROM sessions WHERE token = $1 LIMIT 1)"
	row := api.DB.QueryRow(query, token)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

func TestDeleteSession(t *testing.T) {
	defer clearDbUsers()
	token := "session-token"
	userId := makeDbUser("john.doe@gmail.com", "some-hash")
	makeDbSession(userId, token)
	deleteSession(token)
	if dbSessionExists(token) {
		t.Fail()
	}
}
