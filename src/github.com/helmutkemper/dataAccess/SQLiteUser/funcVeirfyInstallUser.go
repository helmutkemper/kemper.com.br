package SQLiteUser

import (
	"database/sql"
)

func (e *SQLiteUser) verifyInstallUser() (installed bool, err error) {
	var rows *sql.Rows

	rows, err = e.Database.Query(`
		SELECT
      count(*)
		FROM
      sqlite_master
		WHERE
      type=? AND
      name=?;
    `,
		"table",
		"user",
	)
	if err != nil {
		return
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return
		}
	}

	installed = count == 1

	return
}
