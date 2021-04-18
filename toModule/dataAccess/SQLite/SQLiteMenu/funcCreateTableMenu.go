package SQLiteMenu

import (
	"database/sql"
)

//createTableMainMenu (Português): Cria a tabela do menu
func (e *SQLiteMenu) createTableMainMenu() (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS
    	menu (
				id TEXT PRIMARY KEY,
				secondaryId TEXT,   	-- id of parent menu
				typeContent INTEGER,   	-- 0: menu; 1 content
				classroom INTEGER,       -- 0: normal; 1 classroom 
				text TEXT,             	-- menu text
				admin INTEGER,         	-- 0: normal user; 1 admin user
				icon TEXT,             	-- fontawesome icon [optional]
				url TEXT,			   	-- link URL [optional]
				itemOrder INTEGER     	-- menu order
			);
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()

	return
}
