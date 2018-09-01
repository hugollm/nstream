package fetchnotes

import (
	"nstream/data"
	"time"
)

func fetchNotes(userId int, start, end time.Time) []data.Note {
	notes := []data.Note{}
	query := `SELECT id, user_id, content, created_at FROM notes
	WHERE created_at BETWEEN $1 AND $2 AND user_id = $3`
	rows, err := data.DB.Query(query, start, end, userId)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		note := data.Note{}
		rows.Scan(&note.Id, &note.UserId, &note.Content, &note.CreatedAt)
		notes = append(notes, note)
	}
	return notes
}
