package mock

import (
	"fmt"
	"nstream/data"
)

func Clear() {
	_, err := data.DB.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}

func Update(table string, id int, col string, value interface{}) {
	query := fmt.Sprintf("UPDATE %s SET %s = $1 WHERE id = $2", table, col)
	_, err := data.DB.Exec(query, value, id)
	if err != nil {
		panic(err)
	}
}

func Exists(table string, col string, value interface{}) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = $1 LIMIT 1)", table, col)
	row := data.DB.QueryRow(query, value)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}
