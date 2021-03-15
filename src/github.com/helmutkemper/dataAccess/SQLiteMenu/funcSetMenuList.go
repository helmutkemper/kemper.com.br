package SQLiteMenu

import (
	"database/sql"
)

func (e *SQLiteMenu) SetMenuList(menu string, admin int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menuList (menu, admin) VALUES(?, ?)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(menu, admin)
	return
}