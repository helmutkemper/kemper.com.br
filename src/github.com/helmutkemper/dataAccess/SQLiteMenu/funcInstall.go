package SQLiteMenu

func (e *SQLiteMenu) Install() (err error) {
	err = e.truncate()
	if err != nil {
		return
	}

	err = e.createTable()
	if err != nil {
		return
	}

	err = e.populateInitialMenu()
	return
}
