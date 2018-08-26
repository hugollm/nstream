package mock

import "testing"

func TestMockSession(t *testing.T) {
	session := Session()
	if session.Id == 0 || session.UserId == 0 || session.Token == "" || session.CreatedAt.IsZero() {
		t.Fail()
	}
}
