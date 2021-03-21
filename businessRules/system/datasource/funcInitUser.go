package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteUser"
)

func (e *RefList) initUser() (err error) {
	user := SQLiteUser.SQLiteUser{}
	err = user.Connect("./db/database.sqlite")
	if err != nil {
		return
	}

	err = user.Install()
	return
}
