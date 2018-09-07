package updatenote

import (
	"nstream/data"
)

func updateNote(userId int, noteId int, content string) {
	query := "UPDATE notes SET content = $1 WHERE id = $2 AND user_id = $3"
	_, err := data.DB.Exec(query, content, noteId, userId)
	if err != nil {
		panic(err)
	}
}
