package sqlite3

import "database/sql"

func (e *SQLite3) MailExists(mail string) (found bool, err error) {
	var rows *sql.Rows
	rows, err = e.database.Query(
		`
			SELECT 
				id 
			FROM 
				user 
			WHERE 
				mail = ?
		`,
		mail,
	)
	if err != nil {
		return
	}

	found = rows.Next()
	return
}
