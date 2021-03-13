package sqlite3

import "database/sql"

func (e *SQLite3) Connect(connectionString string, args ...interface{}) (err error) {
	e.Database, err = sql.Open("sqlite3", connectionString)
	return
}
