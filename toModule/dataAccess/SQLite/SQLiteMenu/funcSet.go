package SQLiteMenu

import (
	"database/sql"
	"log"
)

// SetMenu (Português): Adiciona um novo ítem ao menu
//func (e *SQLiteMenu) _Set(idMenu, idSecondary string, typeContent int, text string, admin int, icon, url string, order int) (err error) {
//	var statement *sql.Stmt
//	statement, err = e.Database.Prepare(
//		`INSERT INTO main.menu (secondaryId, typeContent, menuId, text, admin, icon, url, itemOrder) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
//	)
//	if err != nil {
//		log.Printf("SQLiteMenu.Set().error: %v", err.Error())
//		return
//	}
//
//	_, err = statement.Exec(idSecondary, typeContent, idMenu, text, admin, icon, url, order)
//	if err != nil {
//		log.Printf("SQLiteMenu.Set().error: %v", err.Error())
//	}
//	return
//}

func (e *SQLiteMenu) Set(id, idMenu, idSecondary string, typeContent int, text string, admin int, icon, url string, order int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menu (id, secondaryId, typeContent, menuId, text, admin, icon, url, itemOrder) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.Printf("SQLiteMenu.SetWithId().error: %v", err.Error())
		return
	}

	_, err = statement.Exec(id, idSecondary, typeContent, idMenu, text, admin, icon, url, order)
	if err != nil {
		log.Printf("SQLiteMenu.SetWithId().error: %v", err.Error())
	}
	return
}
