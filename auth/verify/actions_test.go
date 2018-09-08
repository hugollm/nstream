package verify

import (
	"nstream/data"
	"nstream/data/mock"
	"testing"
)

func TestUserGetsVerified(t *testing.T) {
	user := mock.User()
	token := mock.RandString(64)
	mock.Update("users", user.Id, "verification_token", token)
	verifyUser(token)
	var verified bool
	query := "SELECT verified FROM users WHERE id = $1"
	data.DB.QueryRow(query, user.Id).Scan(&verified)
	if !verified {
		t.Fail()
	}
}
