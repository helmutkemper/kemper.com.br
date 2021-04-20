package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
)

func (e *MongoDBLanguage) CreateTableLanguage() (err error) {

	e.ClientLanguage = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionLanguage)
	return
}
