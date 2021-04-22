package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDB"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBUser struct {
	MongoDB.MongoDB
	ClientUser *mongo.Collection
}
