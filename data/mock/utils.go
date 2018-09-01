package mock

import (
	"encoding/json"
	"fmt"
	"nstream/data"
	"reflect"
)

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

func Json(input interface{}) string {
	bytes, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func HasZeroValues(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Interface() == reflect.Zero(value.Field(i).Type()).Interface() {
			return true
		}
	}
	return false
}
