package sqlite3

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func (e *SQLite3) Connect(connectionString string, args ...interface{}) (err error) {
	e.Database, err = sql.Open("sqlite3", connectionString)
	return
}
