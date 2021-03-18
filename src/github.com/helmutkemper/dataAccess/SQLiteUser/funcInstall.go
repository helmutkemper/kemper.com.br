package SQLiteUser

// Install (Português): Instala o menu e popula os primeiros dados.
func (e *SQLiteUser) Install() (err error) {
	var installed = false

	installed, err = e.verifyInstallUser()
	if err != nil {
		return
	}

	if installed == false {
		err = e.createTableUser()
		if err != nil {
			return
		}

		err = e.populateInitialUser()
		if err != nil {
			return
		}
	}

	return
}
