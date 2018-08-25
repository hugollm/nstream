package login

import (
	"crypto/rand"
	"encoding/hex"
	"nstream/api"
)

type User struct {
	id       int
	email    string
	password string
}

func getUser(qEmail string) (User, error) {
	var id int
	var email, password string
	query := "SELECT id, email, password FROM users WHERE LOWER(email) = LOWER($1) LIMIT 1"
	row := api.DB.QueryRow(query, qEmail)
	err := row.Scan(&id, &email, &password)
	return User{id, email, password}, err
}

func addSession(userId int) (token string) {
	token = makeToken()
	query := "INSERT INTO sessions (user_id, token) VALUES ($1, $2)"
	_, err := api.DB.Exec(query, userId, token)
	if err != nil {
		panic(err)
	}
	return token
}

func makeToken() string {
	buff := make([]byte, 32)
	_, err := rand.Read(buff)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buff)
}
