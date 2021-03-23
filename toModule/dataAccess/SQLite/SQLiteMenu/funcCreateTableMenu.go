package SQLiteMenu

import (
	"database/sql"
)

//createTableMenu (Português): Cria a tabela do menu
func (e *SQLiteMenu) createTableMenu() (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS
    	menu (
				id TEXT PRIMARY KEY,
				menuId TEXT,        -- id menu list
				secondaryId TEXT,   -- id of parent menu
				typeContent INTEGER,   -- 0: menu; 1 content
				text TEXT,             -- menu text
				admin INTEGER,         -- 0: normal user; 1 admin user
				icon TEXT,             -- fontawesome icon [optional]
				url TEXT,			    		 -- link URL [optional]
				itemOrder INTEGER,     -- menu order
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
