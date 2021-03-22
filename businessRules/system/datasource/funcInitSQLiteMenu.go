package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
)

// InitMenu (Português): Inicializa o datasource do menu
func (e *RefList) initSQLiteMenu() (err error) {
	e.Menu = &SQLiteMenu.SQLiteMenu{}
	err = e.Menu.Connect(consts.KDatabaseName)
	if err != nil {
		return
	}

	err = e.Menu.Install()
	return
}
