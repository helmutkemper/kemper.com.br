package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBUser) GetByEmail(mail string) (user dataFormat.User, err error) {
	var cursor *mongo.Cursor
	var userSlice []dataFormat.User
	e.ClientUser = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)

	user = dataFormat.User{Mail: mail}
	cursor, err = e.ClientUser.Find(e.Ctx, user.GetMailAsBSonQuery())
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &userSlice)
	if err != nil {
		util.TraceToLog()
		return
	}

	if len(userSlice) != 0 {
		user = userSlice[0]
	}

	return
}
