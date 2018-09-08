package verify

import (
	"nstream/data/mock"
	"testing"
)

func TestValidInput(t *testing.T) {
	user := mock.User()
	token := mock.RandString(64)
	mock.Update("users", user.Id, "verification_token", token)
	input := jsonInput{token}
	_, errs := validateInput(input)
	if len(errs) > 0 {
		t.Fail()
	}
}

func TestInvalidInput(t *testing.T) {
	input := jsonInput{"invalid-token"}
	_, errs := validateInput(input)
	if len(errs) == 0 || errs["token"].Error() != "Invalid token." {
		t.Fail()
	}
}

func TestValidToken(t *testing.T) {
	user := mock.User()
	token := mock.RandString(64)
	mock.Update("users", user.Id, "verification_token", token)
	vToken, err := validateToken(token)
	if vToken == "" || err != nil {
		t.Fail()
	}
}

func TestInvalidToken(t *testing.T) {
	_, err := validateToken("invalid-token")
	if err == nil || err.Error() != "Invalid token." {
		t.Fail()
	}
}
