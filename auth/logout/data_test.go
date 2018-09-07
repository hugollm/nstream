package logout

import (
	"nstream/data/mock"
	"testing"
)

func TestDeleteSession(t *testing.T) {
	session := mock.Session()
	deleteSession(session.Token)
	if mock.Exists("sessions", "id", session.Id) {
		t.Fail()
	}
}
