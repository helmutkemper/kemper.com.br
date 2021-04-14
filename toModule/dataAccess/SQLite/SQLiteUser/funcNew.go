package SQLiteUser

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"log"
)

func New() (referenceInicialized *SQLiteUser, err error) {
	referenceInicialized = &SQLiteUser{}
	err = referenceInicialized.Connect(consts.KDatabaseName)
	if err != nil {
		log.Printf("SQLiteUser.New().Connect().error: %v", err.Error())
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		log.Printf("SQLiteUser.New().Install().error: %v", err.Error())
		return
	}

	return
}