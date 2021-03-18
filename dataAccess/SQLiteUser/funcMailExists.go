package SQLiteUser

import "database/sql"

func (e *SQLiteUser) MailExists(mail string) (found bool, err error) {
	var rows *sql.Rows
	rows, err = e.Database.Query(
		`
			SELECT 
				id 
			FROM 
				user 
			WHERE 
				eMail = ?
		`,
		mail,
	)
	if err != nil {
		return
	}

	found = rows.Next()
	return
}
