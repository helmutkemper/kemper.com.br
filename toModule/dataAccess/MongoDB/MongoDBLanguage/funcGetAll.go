package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (e *MongoDBLanguage) GetAll() (languagues []dataFormat.Languages, length int, err error) {
	var cursor *mongo.Cursor

	e.ClientLanguage = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionLanguage)
	cursor, err = e.ClientLanguage.Find(e.Ctx, bson.M{})
	if err != nil {
		log.Printf("e.Client.Database().error: %v", err.Error())
		return
	}

	err = cursor.All(e.Ctx, &languagues)
	if err != nil {
		log.Printf("cursor.All().error: %v", err.Error())
		return
	}

	length = len(languagues)
	return
}
