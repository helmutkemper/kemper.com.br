package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"log"
)

func (e *RefList) _initSQLiteUser() (err error) {
	e.User = &SQLiteUser.SQLiteUser{}
	err = e.User.Connect(consts.KDatabaseName)
	if err != nil {
		log.Printf("datasource.initSQLiteUser().error: %v", err.Error())
		return
	}

	err = e.User.Install()
	if err != nil {
		log.Printf("datasource.initSQLiteUser().error: %v", err.Error())
	}

	return
}
