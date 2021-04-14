package SQLiteUser

import (
	"database/sql"
	"log"
)

// Set (Português): Adiciona um novo ítem ao menu
func (e *SQLiteUser) Set(id, idMenu string, admin int, name, nickName, eMail, password string) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.user (id, menuId, admin, name, nickName, eMail, password) VALUES(?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.Printf("SQLiteUser.SetMainMenu().error: %v", err.Error())
		return
	}

	_, err = statement.Exec(id, idMenu, admin, name, nickName, eMail, password)
	if err != nil {
		log.Printf("SQLiteUser.SetMainMenu().error: %v", err.Error())
	}
	return
}
