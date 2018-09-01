package data

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB = makeDB()

func makeDB() *sql.DB {
	connStr := "user=nstream password=nstream dbname=nstream host=/var/run/postgresql"
	db, openErr := sql.Open("postgres", connStr)
	if openErr != nil {
		panic(openErr)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	return db
}
