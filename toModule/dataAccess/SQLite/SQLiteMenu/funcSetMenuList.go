package SQLiteMenu

import (
	"database/sql"
	"log"
)

// SetMenuList (Português): Adiciona um novo ítem ao menu
func (e *SQLiteMenu) _SetMenuList(menu string, admin int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menuList (menu, admin) VALUES(?, ?)`,
	)
	if err != nil {
		log.Printf("SQLiteMenu.SetMenuList().error: %v", err.Error())
		return
	}

	_, err = statement.Exec(menu, admin)
	if err != nil {
		log.Printf("SQLiteMenu.SetMenuList().error: %v", err.Error())
	}
	return
}

//fixme: arquivo
func (e *SQLiteMenu) SetMenuList(id, menu string, admin int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menuList (id, menu, admin) VALUES(?, ?, ?)`,
	)
	if err != nil {
		log.Printf("SQLiteMenu.SetMenuListWithId().error: %v", err.Error())
		return
	}

	_, err = statement.Exec(id, menu, admin)
	if err != nil {
		log.Printf("SQLiteMenu.SetMenuListWithId().error: %v", err.Error())
	}
	return
}
