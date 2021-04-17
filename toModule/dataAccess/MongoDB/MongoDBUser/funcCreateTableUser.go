package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/constants"
)

func (e *MongoDBUser) createTableUser() (err error) {
	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)
	return
}
