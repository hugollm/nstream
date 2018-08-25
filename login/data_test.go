package login

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

func makeDbUser(email string, pass string) {
	_, err := api.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, pass)
	if err != nil {
		panic(err)
	}
}

func getDbSession(tkn string) (id, user_id int, token string) {
	query := "SELECT id, user_id, token FROM sessions WHERE token = $1 LIMIT 1"
	row := api.DB.QueryRow(query, tkn)
	err := row.Scan(&id, &user_id, &token)
	if err != nil {
		panic(err)
	}
	return id, user_id, token
}

func TestGetUserByEmail(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", "some-hash")
	user, err := getUser("john.doe@gmail.com")
	if err != nil || user.id == 0 || user.email != "john.doe@gmail.com" || user.password != "some-hash" {
		t.Fail()
	}
}

func TestGetUserWithUnregisteredEmail(t *testing.T) {
	defer clearDbUsers()
	_, err := getUser("unregistered@gmail.com")
	if err == nil {
		t.Fail()
	}
}

func TestGetUserIsCaseInsensitive(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("JOHN.DOE@gmail.com", "some-hash")
	user, err := getUser("john.doe@GMAIL.COM")
	if err != nil || user.email != "JOHN.DOE@gmail.com" {
		t.Fail()
	}
}

func TestAddSessionReturnsStrongToken(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", "some-hash")
	user, _ := getUser("john.doe@gmail.com")
	token := addSession(user.id)
	if len(token) < 64 {
		t.Fail()
	}
}

func TestAddedSessionGetsPersisted(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", "some-hash")
	user, _ := getUser("john.doe@gmail.com")
	tkn := addSession(user.id)
	id, userId, token := getDbSession(tkn)
	if id == 0 || userId == 0 || token != tkn {
		t.Fail()
	}
}

func TestMakeToken(t *testing.T) {
	token := makeToken()
	if len(token) < 64 {
		t.Fail()
	}
}
