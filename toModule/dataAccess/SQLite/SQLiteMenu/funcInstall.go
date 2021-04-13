package SQLiteMenu

// Install (Português): Instala o menu e popula os primeiros dados.
func (e *SQLiteMenu) Install() (err error) {
	err = e.installMenuList()
	if err != nil {
		return
	}

	err = e.installMainMenu()
	if err != nil {
		return
	}

	return
}

func (e *SQLiteMenu) installMainMenu() (err error) {
	var installed = false

	installed, err = e.verifyInstallMenu()
	if err != nil {
		return
	}

	if installed == false {
		err = e.createTableMenu()
		if err != nil {
			return
		}

		err = e.populateInitialMenu()
		if err != nil {
			return
		}
	}

	return
}

func (e *SQLiteMenu) installMenuList() (err error) {
	var installed = false

	installed, err = e.verifyInstallMenuList()
	if err != nil {
		return
	}

	if installed == false {
		err = e.createTableMenuList()
		if err != nil {
			return
		}

		err = e.populateInitialMenuList()
		if err != nil {
			return
		}
	}

	return
}
