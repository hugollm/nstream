package newnote

import (
	"nstream/data/mock"
	"testing"
)

func TestInsertNote(t *testing.T) {
	user := mock.User()
	note := insertNote(user.Id, "Lorem ipsum.")
	if note.Id == 0 || note.UserId != user.Id || note.Content != "Lorem ipsum." || note.CreatedAt.IsZero() {
		t.Fail()
	}
	if !mock.Exists("notes", "id", note.Id) {
		t.Fail()
	}
}
