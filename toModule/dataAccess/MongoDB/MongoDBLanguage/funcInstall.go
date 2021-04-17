package MongoDBLanguage

import (
	"log"
)

func (e *MongoDBLanguage) Install() (err error) {
	var installed = false

	installed, err = e.verifyInstallLanguage()
	if err != nil {
		log.Printf("MongoDBLanguage.Install().error: %v", err.Error())
		return
	}

	if installed == false {
		err = e.CreateTableLanguage()
		if err != nil {
			log.Printf("MongoDBLanguage.Install().error: %v", err.Error())
			return
		}

		err = e.populateInitialLanguages()
		if err != nil {
			log.Printf("MongoDBLanguage.Install().error: %v", err.Error())
			return
		}
	}

	return
}
