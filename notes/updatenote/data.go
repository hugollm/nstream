package updatenote

import "nstream/data"

func updateNote(id int, content string) {
	_, err := data.DB.Exec("UPDATE notes SET content = $1 WHERE id = $2", content, id)
	if err != nil {
		panic(err)
	}
}
