package main

import (
	"database/sql"
	"log"
)

func (e *SQLiteLanguage) CreateTableLanguage() (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS
    	language (
				id TEXT PRIMARY KEY,
				name TEXT
			);
		`,
	)
	if err != nil {
		log.Printf("SQLiteLanguage.CreateTableLanguage().error: %v", err.Error())
		return
	}

	_, err = statement.Exec()
	if err != nil {
		log.Printf("SQLiteLanguage.CreateTableLanguage().error: %v", err.Error())
	}

	return
}
