package newnote

import (
	"nstream/data"
)

func writeNewNote(userId int, content string) data.Note {
	note := data.Note{
		UserId:  userId,
		Content: content,
	}
	query := "INSERT INTO notes (user_id, content) VALUES ($1, $2) returning id, created_at"
	row := data.DB.QueryRow(query, note.UserId, note.Content)
	err := row.Scan(&note.Id, &note.CreatedAt)
	if err != nil {
		panic(err)
	}
	return note
}
