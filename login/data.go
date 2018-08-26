package login

import (
	"crypto/rand"
	"encoding/hex"
	"nstream/data"
)

func getUser(email string) (data.User, error) {
	var user data.User
	query := "SELECT id, email, password FROM users WHERE LOWER(email) = LOWER($1) LIMIT 1"
	row := data.DB.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	return user, err
}

func addSession(userId int) (token string) {
	token = makeToken()
	query := "INSERT INTO sessions (user_id, token) VALUES ($1, $2)"
	_, err := data.DB.Exec(query, userId, token)
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
