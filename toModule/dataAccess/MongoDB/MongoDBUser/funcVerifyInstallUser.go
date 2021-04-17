package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBUser) verifyInstallUser() (installed bool, err error) {
	var cursor *mongo.Cursor
	var users []dataFormat.User

	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)

	cursor, err = e.ClientUser.Find(e.Ctx, dataFormat.User{Id: constants.KMainUserID, Mail: constants.KMainUserMail})
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &users)
	if err != nil {
		util.TraceToLog()
		return
	}

	installed = len(users) > 0
	return
}
