package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLiteMenu"
)

// initSQLite (Português): Inicializa o SQLite
func (e *RefList) initSQLite() (err error) {
	e.Menu = &SQLiteMenu.SQLiteMenu{}
	err = e.Menu.Connect("./db/database.sqlite")
	return
}
