package SQLiteMenu

import (
	"database/sql"
)

//createTableMenuList (Português): Cria a tabela com todos os menus caso o site tenha múltiplos usuários
func (e *SQLiteMenu) createTableMenuList() (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS
    menuList(
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      menu TEXT,          -- menu name
      admin INTEGER       -- 0: normal user; 1 admin user
    );
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()
	return
}
