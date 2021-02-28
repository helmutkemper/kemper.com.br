package sqlite3

import "database/sql"

func (e *SQLite3) userCreate() (err error) {
	var statement *sql.Stmt
	statement, err = e.database.Prepare(`
		CREATE TABLE IF NOT EXISTS 
			user (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				name TEXT,
				nickname TEXT, 
				gender INTEGER,
				userType INTEGER,
				mail TEXT,
				password TEXT,
				mailVerified INTEGER
			)
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()
	return
}

//                         id, secondaryId, text,                       admin, icon,                       url,                                                   itemOrder
//CREATE TABLE IF NOT EXISTS
//    menu (
//             id INTEGER PRIMARY KEY AUTOINCREMENT,
//             secondaryId INTEGER,   -- id of parent menu
//             text TEXT,             -- menu text
//             admin INTEGER,         -- 0: normal user; 1 admin user
//             icon TEXT,             -- fontawesome icon [optional]
//             url TEXT,			         -- link URL [optional]
//             itemOrder INTEGER      -- menu order
//);
//INSERT INTO main.menu VALUES
//                             (1,  0, 'Kemper.com.br',            0, 'fas fa-code-branch',       '',                                                    0),
//                             (2,  1, 'Sobre mim',                0, 'fas fa-info-circle',       '',                                                    0),
//                             (3,  1, 'Github',                   0, 'fas fa-info-circle',       'https://github.com/helmutkemper',                     1),
//                             (4,  1, 'LinkedIn',                 0, 'fab fa-linkedin',          'https://www.linkedin.com/in/helmut-kemper-93a5441b/', 2),
//                             (5,  0, 'Códigos',                  0, 'fas fa-code',              '',                                                    1),
//                             (6,  5, 'Conversando com o sênior', 0, 'fas fa-fire-extinguisher', '',                                                    2),
//                             (7,  5, 'Migrando para o Go',       0, 'fas fa-fire',              '',                                                    3),
//                             (8,  0, 'Admin',                    0, 'fas fa-cogs',              '',                                                    2),
//                             (9,  8, 'Login',                    0, '',                         '',                                                    1),
//                             (10, 8, 'New content',              0, '',                         '',                                                    3),
//                             (11, 0, 'Donation',                 0, '',                         '',                                                    3);
func (e *SQLite3) menuCreate() (err error) {
	var statement *sql.Stmt
	statement, err = e.database.Prepare(`
		CREATE TABLE IF NOT EXISTS 
			menu (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				secondaryId INTEGER,   -- id of parent menu
				text TEXT,             -- menu text
				admin INTEGER,         -- 0: normal user; 1 admin user
				icon TEXT,             -- fontawesome icon [optional]
				url TEXT,			         -- link URL [optional]
				itemOrder INTEGER      -- menu order
			)
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()
	return
}
