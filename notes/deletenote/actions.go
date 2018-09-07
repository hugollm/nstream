package deletenote

import (
	"nstream/data"
)

func deleteNote(userId, noteId int) {
	_, err := data.DB.Exec("DELETE FROM notes WHERE id = $1 AND user_id = $2", noteId, userId)
	if err != nil {
		panic(err)
	}
}
