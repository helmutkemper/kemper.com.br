package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
)

func (e *RefList) initSQLiteUser() (err error) {
	e.User = &SQLiteUser.SQLiteUser{}
	err = e.User.Connect(consts.KDatabaseName)
	if err != nil {
		return
	}

	err = e.User.Install()
	return
}
