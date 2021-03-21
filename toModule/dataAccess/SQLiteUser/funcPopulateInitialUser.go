package SQLiteUser

// populateInitialMenu (Português): popula o menu com os primeiros dados após a instalação.
func (e *SQLiteUser) populateInitialUser() (err error) {

	err = e.Set(
		1,
		1,
		"Helmut Kemper",
		"Kemper",
		"helmut.kemper@gmail.com",
		"admin",
	)
	return
}
