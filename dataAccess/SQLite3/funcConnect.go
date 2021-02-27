package SQLite3

import "database/sql"

func (e *SQLite3) Connect(filePath string) (err error) {
	e.database, err = sql.Open("sqlite3", filePath)
	return
}
