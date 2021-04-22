package main

import (
	"database/sql"
	"log"
)

func (e *SQLiteLanguage) verifyInstallLanguage() (installed bool, err error) {
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
		"language",
	)
	if err != nil {
		log.Printf("SQLiteLanguage.verifyInstallLanguage().error: %v", err.Error())
		return
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Printf("SQLiteLanguage.verifyInstallLanguage().error: %v", err.Error())
			return
		}
	}

	installed = count == 1

	return
}
