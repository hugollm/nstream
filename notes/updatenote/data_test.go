package updatenote

import (
	"nstream/data"
	"nstream/data/mock"
	"testing"
)

func getNoteContent(noteId int) (content string) {
	data.DB.QueryRow("SELECT content FROM notes WHERE id = $1", noteId).Scan(&content)
	return content
}

func TestUpdateNote(t *testing.T) {
	note := mock.Note()
	updateNote(note.UserId, note.Id, "Lorem ipsum.")
	content := getNoteContent(note.Id)
	if content != "Lorem ipsum." {
		t.Fail()
	}
}

func TestUpdateFailsIfUserIsNotTheOwner(t *testing.T) {
	note := mock.Note()
	updateNote(0, note.Id, "Lorem ipsum.")
	content := getNoteContent(note.Id)
	if content == "Lorem ipsum." {
		t.Fail()
	}
}
