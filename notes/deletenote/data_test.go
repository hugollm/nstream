package deletenote

import (
	"nstream/data/mock"
	"testing"
)

func TestDeleteNote(t *testing.T) {
	note := mock.Note()
	deleteNote(note.UserId, note.Id)
	if mock.Exists("notes", "id", note.Id) {
		t.Fail()
	}
}

func TestDeleteNoteMustFailIfUserIdDontMatch(t *testing.T) {
	note := mock.Note()
	deleteNote(0, note.Id)
	if !mock.Exists("notes", "id", note.Id) {
		t.Fail()
	}
}
