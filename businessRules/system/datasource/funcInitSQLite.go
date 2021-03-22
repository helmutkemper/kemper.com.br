package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
)

//*************************************************************************
//**                               Cuidado                               **
//*************************************************************************

// initSQLite (Português): Inicializa o SQLite
func (e *RefList) initSQLite() (err error) {
	e.Menu = &SQLiteMenu.SQLiteMenu{}
	err = e.Menu.Connect(consts.KDatabaseName)

	e.User = &SQLiteUser.SQLiteUser{}
	err = e.User.Connect(consts.KDatabaseName)
	return
}
