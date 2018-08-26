package data

import "time"

type User struct {
	Id        int
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	UserId    int
	Token     string
	CreatedAt time.Time
}
