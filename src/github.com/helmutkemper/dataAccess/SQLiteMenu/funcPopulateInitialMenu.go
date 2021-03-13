package SQLiteMenu

func (e *SQLiteMenu) populateInitialMenu() (err error) {

	err = e.SetMenu(1, 0, "Kemper.com.br", 0, "fas fa-code-branch", "", 0)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 0, "Códigos", 0, "fas fa-code", "", 1)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 0, "Admin", 1, "fas fa-cogs", "", 2)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 0, "Donation", 0, "", "", 3)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 1, "Sobre mim", 0, "fas fa-info-circle", "", 0)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 1, "Github", 0, "fas fa-info-circle", "https://github.com/helmutkemper", 1)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 1, "LinkedIn", 0, "fab fa-linkedin", "https://www.linkedin.com/in/helmut-kemper-93a5441b/", 2)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 2, "Conversando com o sênior", 0, "fas fa-fire-extinguisher", "", 0)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 2, "Migrando para o Go", 0, "fas fa-fire", "", 1)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 3, "Login", 1, "", "", 1)
	if err != nil {
		return
	}

	err = e.SetMenu(1, 3, "New content", 1, "", "", 0)
	return
}
