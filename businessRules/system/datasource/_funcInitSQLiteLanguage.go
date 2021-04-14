package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteLanguage"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"log"
)

func (e *RefList) _initSQLiteLanguage() (err error) {
	e.Language = &SQLiteLanguage.SQLiteLanguage{}
	err = e.Language.Connect(consts.KDatabaseName)
	if err != nil {
		log.Printf("datasource.initSQLiteLanguage().error: %v", err.Error())
		return
	}

	err = e.Language.Install()
	if err != nil {
		log.Printf("datasource.initSQLiteLanguage().error: %v", err.Error())
	}

	return
}
