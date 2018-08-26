package mock

import "nstream/data"

func Session() data.Session {
	user := User()
	session := data.Session{
		UserId: user.Id,
		Token:  RandString(64),
	}
	query := "INSERT INTO sessions (user_id, token) VALUES ($1, $2) RETURNING id, created_at"
	row := data.DB.QueryRow(query, session.UserId, session.Token)
	err := row.Scan(&session.Id, &session.CreatedAt)
	if err != nil {
		panic(err)
	}
	return session
}
