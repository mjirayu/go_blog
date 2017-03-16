package models

import (
	"database/sql"

	// PG adapter
	_ "github.com/lib/pq"
)

// Db : Database Connection
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/go_blog?sslmode=disable")
	if err != nil {
		panic(err)
	}
}
