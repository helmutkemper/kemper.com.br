package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
)

func (e *RefList) initUser() (err error) {
	user := SQLiteUser.SQLiteUser{}
	err = user.Connect(consts.KDatabaseName)
	if err != nil {
		return
	}

	err = user.Install()
	return
}
