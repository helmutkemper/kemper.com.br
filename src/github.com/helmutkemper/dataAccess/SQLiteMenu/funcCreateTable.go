package SQLiteMenu

import (
	"database/sql"
)

func (e *SQLiteMenu) createTable() (err error) {
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
	if err != nil {
		return
	}

	statement, err = e.Database.Prepare(`
		INSERT INTO main.menuList VALUES (1, 'Menu Principal', 0);
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()
	if err != nil {
		return
	}

	statement, err = e.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS
    	menu (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				menuId INTEGER,        -- id menu list
				secondaryId INTEGER,   -- id of parent menu
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
