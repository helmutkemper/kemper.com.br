package SQLiteUser

import (
	"database/sql"
)

// SetMenu (Português): Adiciona um novo ítem ao menu
func (e *SQLiteUser) Set(idMenu, admin int, name, nickName, eMail, password string) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.user (menuId, admin, name, nickName, eMail, password) VALUES(?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(idMenu, admin, idMenu, name, nickName, eMail, password)
	return
}
