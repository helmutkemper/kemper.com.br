package SQLiteMenu

import (
	"database/sql"
)

func (e *SQLiteMenu) truncate() (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`DROP TABLE IF EXISTS main.menu;`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()
	return
}
