package fetchnotes

import (
	"nstream/data/mock"
	"testing"
	"time"
)

func TestFetchNotesIsRestrictedToTheSpecifiedUser(t *testing.T) {
	user := mock.User()
	note := mock.Note()
	mock.Update("notes", note.Id, "user_id", user.Id)
	mock.Note() // should not be fetched
	now := time.Now()
	lastWeek := now.AddDate(0, 0, -7)
	notes := fetchNotes(user.Id, lastWeek, now)
	if len(notes) != 1 {
		t.FailNow()
	}
	if mock.HasZeroValues(notes[0]) {
		t.Fail()
	}
}
