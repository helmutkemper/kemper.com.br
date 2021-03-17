package SQLiteMenu

import (
	"database/sql"
)

// SetMenu (Português): Adiciona um novo ítem ao menu
func (e *SQLiteMenu) SetMenu(idMenu, idSecondary, typeContent int, text string, admin int, icon, url string, order int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menu (secondaryId, typeContent, menuId, text, admin, icon, url, itemOrder) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(idSecondary, typeContent, idMenu, text, admin, icon, url, order)
	return
}
