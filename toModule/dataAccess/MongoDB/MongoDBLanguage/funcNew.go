package MongoDBLanguage

import (
	"log"
)

func New() (referenceInicialized *MongoDBLanguage, err error) {
	referenceInicialized = &MongoDBLanguage{}
	err = referenceInicialized.Connect("mongodb://127.0.0.1:27017/")
	if err != nil {
		log.Printf("MongoDBLanguage.New().Connect().error: %v", err.Error())
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		log.Printf("MongoDBLanguage.New().Install().error: %v", err.Error())
		return
	}

	return
}
