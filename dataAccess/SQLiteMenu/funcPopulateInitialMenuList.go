package SQLiteMenu

// populateInitialMenuList (Português): popula a lista de menu após a instalação
func (e *SQLiteMenu) populateInitialMenuList() (err error) {

	err = e.SetMenuList("Menu Principal", 0)
	if err != nil {
		return
	}

	return
}
