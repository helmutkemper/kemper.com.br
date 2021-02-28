package sqlite3

import "database/sql"

func (e *SQLite3) UserInsert(name, nickname string, gender, userType int, mail, password string, mailVerified int) (err error) {
	var statement *sql.Stmt
	statement, err = e.database.Prepare(`
		INSERT INTO 
		    user (name, nickname, gender, userType, mail, password, mailVerified) 
		VALUES   
		       	 (?,    ?,        ?,      ?,        ?,    ?,        ?)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(name, nickname, gender, userType, mail, password, mailVerified)
	return
}
