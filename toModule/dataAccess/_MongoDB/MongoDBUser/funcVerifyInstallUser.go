package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBUser) verifyInstallUser() (installed bool, err error) {
	var cursor *mongo.Cursor
	var users []dataFormat.User
	var user = dataFormat.User{Id: constants.KMainUserID, Mail: constants.KMainUserMail}

	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)

	cursor, err = e.ClientUser.Find(e.Ctx, user.GetIdAndMailAsBSonQuery())
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
