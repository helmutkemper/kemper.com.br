package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
)

// InitMenu (Português): Inicializa o datasource do menu
func (e *RefList) initMenu() (err error) {
	menu := SQLiteMenu.SQLiteMenu{}
	err = menu.Connect(consts.KDatabaseName)
	if err != nil {
		return
	}

	err = menu.Install()
	return
}
