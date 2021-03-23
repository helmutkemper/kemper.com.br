package SQLiteUser

import (
	"log"
)

// Install (Português): Instala o menu e popula os primeiros dados.
func (e *SQLiteUser) Install() (err error) {
	var installed = false

	installed, err = e.verifyInstallUser()
	if err != nil {
		log.Printf("SQLiteUser.Install().error: %v", err.Error())
		return
	}

	if installed == false {
		err = e.createTableUser()
		if err != nil {
			log.Printf("SQLiteUser.Install().error: %v", err.Error())
			return
		}

		err = e.populateInitialUser()
		if err != nil {
			log.Printf("SQLiteUser.Install().error: %v", err.Error())
			return
		}
	}

	return
}
