package MongoDBLanguage

import (
	"github.com/helmutkemper/kemper.com.br/util"
)

func New() (referenceInicialized *MongoDBLanguage, err error) {
	referenceInicialized = &MongoDBLanguage{}
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
