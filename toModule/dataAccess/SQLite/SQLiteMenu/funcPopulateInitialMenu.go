package SQLiteMenu

import (
	"log"
)

// populateInitialMenu (Português): popula o menu com os primeiros dados após a instalação.
func (e *SQLiteMenu) populateInitialMenu() (err error) {

	err = e.SetMainMenu("04daa10f-1059-446a-a2b9-b48d536c8b23", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "", 0, 0, "Kemper.com.br", 0, "fas fa-code-branch", "", 0)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("0409921a-600b-49db-b8cc-d906e3837082", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "", 0, 0, "Códigos", 0, "fas fa-code", "", 1)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("1726ce36-7338-4932-858b-6863700dbe51", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "", 0, 0, "Admin", 1, "fas fa-cogs", "", 2)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("d6667be4-6794-4c6f-9eb2-0d3cac1d7801", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "", 0, 0, "Donation", 0, "fas fa-donate", "", 3)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("d791c1a2-fc36-454a-a7f4-038149e30e4a", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "04daa10f-1059-446a-a2b9-b48d536c8b23", 1, 0, "Sobre mim", 0, "fas fa-info-circle", "", 0)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("d40894ba-8015-402f-9c10-5a4836413b1d", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "04daa10f-1059-446a-a2b9-b48d536c8b23", 1, 0, "Github", 0, "fab fa-github", "https://github.com/helmutkemper", 1)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("873da901-8285-4846-85df-59f16850429e", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "04daa10f-1059-446a-a2b9-b48d536c8b23", 1, 0, "LinkedIn", 0, "fab fa-linkedin", "https://www.linkedin.com/in/helmut-kemper-93a5441b/", 2)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("207b66f1-eef9-4572-a57d-5152b511e37d", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "0409921a-600b-49db-b8cc-d906e3837082", 0, 1, "Conversando com o sênior", 0, "fas fa-fire-extinguisher", "", 0)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("ae116cdf-1941-4ca7-9129-0bad2793ca42", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "0409921a-600b-49db-b8cc-d906e3837082", 0, 1, "Migrando para o Go", 0, "fas fa-fire", "", 1)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("08c7e064-8225-4ee4-9cf4-19dcdd546c84", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "1726ce36-7338-4932-858b-6863700dbe51", 1, 0, "Login", 1, "", "", 1)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu("24fda64b-0af2-4da6-9d95-38664614d3bd", "5996b891-9d3c-4038-af37-cb07f5f0f72d", "1726ce36-7338-4932-858b-6863700dbe51", 1, 0, "New content", 1, "", "", 0)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
	}

	return
}
