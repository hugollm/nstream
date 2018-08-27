package mock

import (
	"nstream/data"
)

func Note() data.Note {
	user := User()
	note := data.Note{
		UserId:  user.Id,
		Content: RandString(200),
	}
	query := "INSERT INTO notes (user_id, content) VALUES ($1, $2) RETURNING id, created_at"
	row := data.DB.QueryRow(query, note.UserId, note.Content)
	err := row.Scan(&note.Id, &note.CreatedAt)
	if err != nil {
		panic(err)
	}
	return note
}
