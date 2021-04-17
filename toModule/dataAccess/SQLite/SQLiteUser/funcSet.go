package SQLiteUser

import (
	"database/sql"
	"log"
)

// Set (Português): Adiciona um novo usuário
func (e *SQLiteUser) Set(id, idMenu string, admin int, name, nickName, email, password string) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.user (id, menuId, admin, name, nickName, email, password) VALUES(?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		log.Printf("SQLiteUser.SetMainMenu().error: %v", err.Error())
		return
	}

	_, err = statement.Exec(id, idMenu, admin, name, nickName, email, password)
	if err != nil {
		log.Printf("SQLiteUser.SetMainMenu().error: %v", err.Error())
	}
	return
}
