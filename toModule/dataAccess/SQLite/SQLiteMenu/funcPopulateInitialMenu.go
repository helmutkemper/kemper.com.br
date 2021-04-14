package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"log"
)

// populateInitialMenu (Português): popula o menu com os primeiros dados após a instalação.
func (e *SQLiteMenu) populateInitialMenu() (err error) {

	err = e.SetMainMenu(
		constants.KMainMenuSiteNameID,
		constants.KMainMenuID,
		"",
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuSiteNameText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuSiteNameIcon,
		constants.KMainMenuSiteNameURL,
		constants.KMainMenuSiteNameOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		constants.KMainMenuCodeID,
		constants.KMainMenuID,
		"",
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuCodeText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuCodeIcon,
		constants.KMainMenuCodeURL,
		constants.KMainMenuCodeOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		constants.KMainMenuAdminID,
		constants.KMainMenuID,
		"",
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuAdminText,
		constants.KmainMenuUserAdmin,
		constants.KMainMenuAdminIcon,
		constants.KMainMenuAdminURL,
		constants.KMainMenuAdminOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		constants.KMainMenuDonationID,
		constants.KMainMenuID,
		"",
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuDonationText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuDonationIcon,
		constants.KMainMenuDonationURL,
		constants.KMainMenuDonationOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		"d791c1a2-fc36-454a-a7f4-038149e30e4a",
		constants.KMainMenuID,
		constants.KMainMenuSiteNameID,
		constants.KMainMenuTypeContentAsContent,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuAboutMeText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuAboutMeIcon,
		constants.KMainMenuAboutMeURL,
		constants.KMainMenuAboutMeOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		"d40894ba-8015-402f-9c10-5a4836413b1d",
		constants.KMainMenuID,
		constants.KMainMenuSiteNameID,
		constants.KMainMenuTypeContentAsContent,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuGithubText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuGithubIcon,
		constants.KMainMenuGithubURL,
		constants.KMainMenuGithubOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		"873da901-8285-4846-85df-59f16850429e",
		constants.KMainMenuID,
		constants.KMainMenuSiteNameID,
		constants.KMainMenuTypeContentAsContent,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuLinkedinText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuLinkedinIcon,
		constants.KMainMenuLinkedinURL,
		constants.KMainMenuLinkedinOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		"207b66f1-eef9-4572-a57d-5152b511e37d",
		constants.KMainMenuID,
		constants.KMainMenuCodeID,
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsClass,
		"Conversando com o sênior",
		constants.KmainMenuUserNormal,
		"fas fa-fire-extinguisher",
		"",
		0,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		"ae116cdf-1941-4ca7-9129-0bad2793ca42",
		constants.KMainMenuID,
		constants.KMainMenuCodeID,
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsClass,
		"Migrando para o Go",
		constants.KmainMenuUserNormal,
		"fas fa-fire",
		"",
		1,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		"08c7e064-8225-4ee4-9cf4-19dcdd546c84",
		constants.KMainMenuID,
		constants.KMainMenuAdminID,
		constants.KMainMenuTypeContentAsContent,
		constants.KmainMenuTypeClassRoomAsNormal,
		"Login",
		constants.KmainMenuUserAdmin,
		"",
		"",
		1,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		"24fda64b-0af2-4da6-9d95-38664614d3bd",
		constants.KMainMenuID,
		constants.KMainMenuAdminID,
		constants.KMainMenuTypeContentAsContent,
		constants.KmainMenuTypeClassRoomAsNormal,
		"New content",
		constants.KmainMenuUserAdmin,
		"",
		"",
		0,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
	}

	return
}
