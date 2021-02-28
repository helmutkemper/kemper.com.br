package sqlite3

import (
	"database/sql"
)

func (e *SQLite3) SetMenu(idSecondary int, text string, admin int, icon, url string, order int) (err error) {
	var statement *sql.Stmt
	statement, err = e.database.Prepare(
		`INSERT INTO main.menu (secondaryId, text, admin, icon, url, itemOrder) VALUES(?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(idSecondary, text, admin, icon, url, order)
	return
}
