package SQLite3

import "database/sql"

func (e *SQLite3) MailHasVerified(mail string) (mailHasVerified bool, err error) {
	var rows *sql.Rows
	rows, err = e.database.Query(
		`
			SELECT 
				   id 
			FROM 
				 user 
			WHERE 
				  mail = ? 
			  AND mailVerified = 1
		  `,
		mail,
	)
	if err != nil {
		return
	}

	mailHasVerified = rows.Next()
	return
}
