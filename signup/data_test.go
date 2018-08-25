package signup

import (
	"nstream/api"
	"testing"
)

func clearDbUsers() {
	api.DB.Exec("DELETE FROM users")
}

func makeDbUser(email string, pass string) {
	api.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, pass)
}

func getDbUser(qEmail string) (id int, email string, password string) {
	query := "SELECT id, email, password FROM users WHERE email = $1 LIMIT 1"
	row := api.DB.QueryRow(query, qEmail)
	row.Scan(&id, &email, &password)
	return
}

func TestUserWithEmailExistsWithExactEmail(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("john.doe@gmail.com", "some-hash")
	if userWithEmailExists("john.doe@gmail.com") != true {
		t.Fail()
	}
}

func TestUserWithEmailExistsIsCaseInsensitive(t *testing.T) {
	defer clearDbUsers()
	makeDbUser("john.doe@GMAIL.COM", "some-hash")
	if userWithEmailExists("JOHN.DOE@gmail.com") != true {
		t.Fail()
	}
}

func TestAddUserPersistsOnDb(t *testing.T) {
	defer clearDbUsers()
	addUser("john.doe@gmail.com", "some-hash")
	id, email, password := getDbUser("john.doe@gmail.com")
	if id <= 0 || email != "john.doe@gmail.com" || password != "some-hash" {
		t.Fail()
	}
}
