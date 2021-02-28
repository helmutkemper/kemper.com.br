package sqlite3

import "database/sql"

type SQLite3 struct {
	database *sql.DB
}
