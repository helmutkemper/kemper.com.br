package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/util"
)

func (e *MongoDBLanguage) Install() (err error) {
	var installed = false

	e.ClientLanguage = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionLanguage)

	installed, err = e.verifyInstallLanguage()
	if err != nil {
		util.TraceToLog()
		return
	}

	if installed == false {
		err = e.CreateTableLanguage()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.populateInitialLanguages()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}
