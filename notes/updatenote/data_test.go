package updatenote

import (
	"testing"
	"nstream/data/mock"
	"nstream/data"
)

func TestUpdateNote(t *testing.T) {
	note := mock.Note()
	updateNote(note.Id, "Lorem ipsum.")
	var content string
	data.DB.QueryRow("SELECT content FROM notes WHERE id = $1", note.Id).Scan(&content)
	if content != "Lorem ipsum." {
		t.Fail()
	}
}
