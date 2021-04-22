package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
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
