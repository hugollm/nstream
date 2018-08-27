package new

import (
	"nstream/data/mock"
	"testing"
)

func TestWriteNewNote(t *testing.T) {
	user := mock.User()
	note := writeNewNote(user.Id, "Lorem ipsum.")
	if note.Id == 0 || note.UserId != user.Id || note.Content != "Lorem ipsum." || note.CreatedAt.IsZero() {
		t.Fail()
	}
	if !mock.Exists("notes", "id", note.Id) {
		t.Fail()
	}
}
