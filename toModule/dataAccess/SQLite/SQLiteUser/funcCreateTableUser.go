package SQLiteUser

import (
	"database/sql"
)

//createTableMenu (Português): Cria a tabela do menu
func (e *SQLiteUser) createTableUser() (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS
    	user (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				menuId INTEGER,        -- id menu list
				admin INTEGER,         -- 0: normal user; 1 admin user
				name TEXT,						 -- complete name
				nickName TEXT,				 -- nick name
				eMail TEXT,						 -- e-mail
				password TEXT,				 -- password
				FOREIGN KEY(menuId) REFERENCES menuList(id) ON DELETE restrict ON UPDATE restrict
			);
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()

	return
}
