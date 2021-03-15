package SQLiteMenu

func (e *SQLiteMenu) populateInitialMenuList() (err error) {

	err = e.SetMenuList("Menu Principal", 0)
	if err != nil {
		return
	}

	return
}
