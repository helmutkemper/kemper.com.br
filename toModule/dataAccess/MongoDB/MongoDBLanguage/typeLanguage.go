package MongoDBLanguage

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDB"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBLanguage struct {
	MongoDB.MongoDB
	ClientLanguage *mongo.Collection
}
