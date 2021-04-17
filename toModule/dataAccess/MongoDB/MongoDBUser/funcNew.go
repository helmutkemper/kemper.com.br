package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/util"
)

func New() (referenceInicialized *MongoDBUser, err error) {
	referenceInicialized = &MongoDBUser{}
	err = referenceInicialized.Connect("mongodb://127.0.0.1:27017/")
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
