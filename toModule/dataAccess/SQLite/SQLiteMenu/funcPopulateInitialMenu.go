package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/util"
)

// populateInitialMenu (Português): popula o menu com os primeiros dados após a instalação.
func (e *SQLiteMenu) populateInitialMenu() (err error) {

	err = e.Set(
		constants.KMainMenuSiteNameID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuCodeID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuAdminID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuDonationID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuAboutMeID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuGithubID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuLinkedinID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuTalkingWithDevID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuMigratingToGoID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuLoginID,
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
		util.TraceToLog()
		return
	}

	err = e.Set(
		constants.KMainMenuNewContentID,
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
		util.TraceToLog()
		return
	}

	return
}
