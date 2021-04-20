package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/util"
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
