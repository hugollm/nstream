package mock

import (
	"testing"
)

func TestMockNote(t *testing.T) {
	note := Note()
	if note.Id == 0 || note.UserId == 0 || note.Content == "" || note.CreatedAt.IsZero() {
		t.Fail()
	}
}
