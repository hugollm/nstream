package mock

import "testing"

func TestMockUser(t *testing.T) {
	user := User()
	if user.Id == 0 || user.Email == "" || user.Password == "" || user.CreatedAt.IsZero() {
		t.Fail()
	}
}
