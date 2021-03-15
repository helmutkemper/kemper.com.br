package SQLiteMenu

func (e *SQLiteMenu) Install() (err error) {
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
