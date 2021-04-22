package main

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDB"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoLanguage MongoDBLanguage

type MongoDBLanguage struct {
	MongoDB.MongoDB
	ClientLanguage *mongo.Collection
}
