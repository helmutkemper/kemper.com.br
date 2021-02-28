package sqlite3

import (
	"database/sql"
)

func (e *SQLite3) GetPassword(mail string) (password string, err error) {
	var rows *sql.Rows
	rows, err = e.database.Query(`
		SELECT 
		       password 
		FROM 
		     user 
		WHERE 
		      mail = ?`,
		mail,
	)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&password)
		return
	}

	return
}
