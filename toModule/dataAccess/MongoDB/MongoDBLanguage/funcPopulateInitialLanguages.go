package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
)

func (e *MongoDBLanguage) populateInitialLanguages() (err error) {
	err = e.Set(constants.KInstallLanguageID, constants.KInstallLanguageName)
	return
}
