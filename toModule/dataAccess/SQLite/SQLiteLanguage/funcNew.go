package SQLiteLanguage

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"log"
)

func New() (referenceInicialized *SQLiteLanguage, err error) {
	referenceInicialized = &SQLiteLanguage{}
	err = referenceInicialized.Connect(consts.KDatabaseName)
	if err != nil {
		log.Printf("SQLiteLanguage.New().Connect().error: %v", err.Error())
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		log.Printf("SQLiteLanguage.New().Install().error: %v", err.Error())
		return
	}

	return
}
