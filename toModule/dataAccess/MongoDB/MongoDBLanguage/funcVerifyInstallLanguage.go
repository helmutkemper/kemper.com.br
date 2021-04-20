package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBLanguage) verifyInstallLanguage() (installed bool, err error) {
	var cursor *mongo.Cursor
	var languagues []dataFormat.Languages

	//fixme: toBSon()
	cursor, err = e.ClientLanguage.Find(e.Ctx, bson.M{"_id": constants.KInstallLanguageID, "name": constants.KInstallLanguageName})
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &languagues)
	if err != nil {
		util.TraceToLog()
		return
	}

	installed = len(languagues) > 0
	return
}
