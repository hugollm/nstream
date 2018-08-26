package mock

import "nstream/data"

func User() data.User {
	user := data.User{
		Email:    RandString(200) + "@gmail.com",
		Password: RandString(60),
	}
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, created_at"
	row := data.DB.QueryRow(query, user.Email, user.Password)
	err := row.Scan(&user.Id, &user.CreatedAt)
	if err != nil {
		panic(err)
	}
	return user
}
