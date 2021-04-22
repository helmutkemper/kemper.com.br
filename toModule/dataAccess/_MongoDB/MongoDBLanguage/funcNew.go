package main

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func New() (referenceInicialized *MongoDBLanguage, err error) {
	referenceInicialized = &MongoDBLanguage{}
	err = referenceInicialized.Connect(constants.KMongoDBConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
