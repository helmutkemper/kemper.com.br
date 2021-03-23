package SQLiteUser

import (
	"database/sql"
	"log"
)

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
		log.Printf("SQLiteUser.MailExists().error: %v", err.Error())
		return
	}

	found = rows.Next()
	return
}
