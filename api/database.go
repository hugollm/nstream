package api

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB = makeDB()

func makeDB() *sql.DB {
	connStr := "user=postgres dbname=nstream host=/var/run/postgresql"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
