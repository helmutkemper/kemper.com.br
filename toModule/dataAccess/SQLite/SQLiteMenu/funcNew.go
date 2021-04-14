package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"log"
)

func New() (referenceInicialized *SQLiteMenu, err error) {
	referenceInicialized = &SQLiteMenu{}
	err = referenceInicialized.Connect(consts.KDatabaseName)
	if err != nil {
		log.Printf("SQLiteMenu.New().Connect().error: %v", err.Error())
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		log.Printf("SQLiteMenu.New().Install().error: %v", err.Error())
		return
	}

	return
}