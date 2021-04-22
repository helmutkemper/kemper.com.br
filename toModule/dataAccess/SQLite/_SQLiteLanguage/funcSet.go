package main

import (
	"database/sql"
	"log"
)

func (e *SQLiteLanguage) Set(id, name string) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.language (id, name) VALUES(?, ?)`,
	)
	if err != nil {
		log.Printf("SQLiteLanguage.Set().error: %v", err.Error())
		return
	}

	_, err = statement.Exec(id, name)
	if err != nil {
		log.Printf("SQLiteLanguage.Set().error: %v", err.Error())
	}
	return
}
