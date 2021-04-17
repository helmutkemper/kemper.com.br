package MongoDBLanguage

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (e *MongoDBLanguage) verifyInstallLanguage() (installed bool, err error) {
	var cursor *mongo.Cursor
	cursor, err = e.ClientLanguage.Find(e.Ctx, bson.M{"_id": constants.KInstallLanguageID, "name": constants.KInstallLanguageName})
	if err != nil {
		log.Printf("e.ClientLanguage.Find().error: %v", err.Error())
		return
	}

	var er = cursor.Err()
	if er != nil {
		return
	}

	installed = true
	return
}
