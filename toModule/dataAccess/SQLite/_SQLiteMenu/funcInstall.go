package SQLiteMenu

import (
	"github.com/helmutkemper/util"
)

// Install (Português): Instala o menu e popula os primeiros dados.
func (e *SQLiteMenu) Install() (err error) {
	var installed = false

	installed, err = e.verifyInstallMenu()
	if err != nil {
		util.TraceToLog()
		return
	}

	if installed == false {
		err = e.createTableMainMenu()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.populateInitialMenu()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}
