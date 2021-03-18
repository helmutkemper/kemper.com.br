package SQLiteUser

import (
	"database/sql"
)

func (e *SQLiteUser) GetPassword(mail string) (password string, err error) {
	var rows *sql.Rows
	rows, err = e.Database.Query(`
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
