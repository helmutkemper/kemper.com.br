package main

import (
	sqlite3 "github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLite3"
)

var Language SQLiteLanguage

type SQLiteLanguage struct {
	sqlite3.SQLite3
}
