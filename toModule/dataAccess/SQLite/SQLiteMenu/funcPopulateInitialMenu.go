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
		constants.KMainMenuAboutMeID,
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
		constants.KMainMenuGithubID,
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
		constants.KMainMenuLinkedinID,
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
		constants.KMainMenuTalkingWithDevID,
		constants.KMainMenuID,
		constants.KMainMenuCodeID,
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsClass,
		constants.KMainMenuTalkingWithDevText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuTalkingWithDevIcon,
		constants.KMainMenuTalkingWithDevURL,
		constants.KMainMenuTalkingWithDevOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		constants.KMainMenuMigratingToGoID,
		constants.KMainMenuID,
		constants.KMainMenuCodeID,
		constants.KMainMenuTypeContentAsMenu,
		constants.KmainMenuTypeClassRoomAsClass,
		constants.KMainMenuMigratingToGoText,
		constants.KmainMenuUserNormal,
		constants.KMainMenuMigratingToGoIcon,
		constants.KMainMenuMigratingToGoURL,
		constants.KMainMenuMigratingToGoOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		constants.KMainMenuLoginID,
		constants.KMainMenuID,
		constants.KMainMenuAdminID,
		constants.KMainMenuTypeContentAsContent,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuLoginText,
		constants.KmainMenuUserAdmin,
		constants.KMainMenuLoginIcon,
		constants.KMainMenuLoginURL,
		constants.KMainMenuLoginOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
		return
	}

	err = e.SetMainMenu(
		constants.KMainMenuNewContentID,
		constants.KMainMenuID,
		constants.KMainMenuAdminID,
		constants.KMainMenuTypeContentAsContent,
		constants.KmainMenuTypeClassRoomAsNormal,
		constants.KMainMenuNewContentText,
		constants.KmainMenuUserAdmin,
		constants.KMainMenuNewContentIcon,
		constants.KMainMenuNewContentURL,
		constants.KMainMenuNewContentOrder,
	)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenu().error: %v", err.Error())
	}

	return
}
