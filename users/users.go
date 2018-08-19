package users

import "errors"

type User struct {
    Id int
    Email string
    Password string
}

var users []User = []User{}

func AddUser(email string, password string) {
    user := User{Email: email, Password: password}
    users = append(users, user)
}

func GetUserByEmail(email string) (User, error) {
    for _, user := range users {
        if user.Email == email {
            return user, nil
        }
    }
    return User{}, errors.New("User not found.")
}
