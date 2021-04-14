package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"log"
)

// InitMenu (Português): Inicializa o datasource do menu
func (e *RefList) _initSQLiteMenu() (err error) {
	e.Menu = &SQLiteMenu.SQLiteMenu{}
	err = e.Menu.Connect(consts.KDatabaseName)
	if err != nil {
		log.Printf("datasource.initSQLiteMenu().error: %v", err.Error())
		return
	}

	err = e.Menu.Install()
	if err != nil {
		log.Printf("datasource.initSQLiteMenu().error: %v", err.Error())
	}
	return
}
