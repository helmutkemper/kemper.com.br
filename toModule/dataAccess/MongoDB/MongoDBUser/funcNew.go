package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/util"
)

func New() (referenceInicialized *MongoDBUser, err error) {
	referenceInicialized = &MongoDBUser{}
	err = referenceInicialized.Connect(constants.KMongoDBConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
