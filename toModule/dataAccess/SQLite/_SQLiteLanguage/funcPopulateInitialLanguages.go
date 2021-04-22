package main

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
)

func (e *SQLiteLanguage) populateInitialLanguages() (err error) {
	err = e.Set(constants.KInstallLanguageID, constants.KInstallLanguageName)
	return
}
