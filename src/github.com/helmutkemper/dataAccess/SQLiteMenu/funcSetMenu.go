package SQLiteMenu

import (
	"database/sql"
)

func (e *SQLiteMenu) SetMenu(idSecondary, idMenu int, text string, admin int, icon, url string, order int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menu (secondaryId, menuId, text, admin, icon, url, itemOrder) VALUES(?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(idSecondary, idMenu, text, admin, icon, url, order)
	return
}
