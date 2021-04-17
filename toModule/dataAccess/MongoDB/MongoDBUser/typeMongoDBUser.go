package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDB"
	"github.com/helmutkemper/kemper.com.br/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBUser struct {
	MongoDB.MongoDB
	ClientUser *mongo.Collection
}

func (e *MongoDBUser) MailExists(mail string) (found bool, err error) {
	var user dataFormat.User
	user, err = e.GetByEmail(mail)
	if err != nil {
		util.TraceToLog()
		return
	}

	found = user.Mail != ""
	return
}
