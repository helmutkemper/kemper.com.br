package sqlite3

import (
	"database/sql"
)

type SQLite3 struct {
	Database *sql.DB
}
