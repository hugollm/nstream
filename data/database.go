package data

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB = makeDB()

func makeDB() *sql.DB {
	connStr := "user=postgres dbname=nstream host=/var/run/postgresql"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
