package SQLite3

import "database/sql"

func (e *SQLite3) userCreate() (err error) {
	var statement *sql.Stmt
	statement, err = e.database.Prepare(`
		CREATE TABLE IF NOT EXISTS 
			user (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				name TEXT,
				nickname TEXT, 
				gender INTEGER,
				userType INTEGER,
				mail TEXT,
				password TEXT,
				mailVerified INTEGER
			)
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()
	return
}
