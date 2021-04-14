package SQLiteMenu

import (
	"database/sql"
	"log"
)

func (e *SQLiteMenu) SetMainMenu(id, idMenu, idSecondary string, typeContent int, classroom int, text string, admin int, icon, url string, order int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menu (id, secondaryId, typeContent, classroom, menuId, text, admin, icon, url, itemOrder) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.Printf("SQLiteMenu.SetWithId().error: %v", err.Error())
		return
	}

	_, err = statement.Exec(id, idSecondary, typeContent, classroom, idMenu, text, admin, icon, url, order)
	if err != nil {
		log.Printf("SQLiteMenu.SetWithId().error: %v", err.Error())
	}
	return
}
