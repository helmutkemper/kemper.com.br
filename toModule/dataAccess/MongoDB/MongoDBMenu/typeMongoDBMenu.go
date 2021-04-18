package MongoDBMenu

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDB"
	"github.com/helmutkemper/kemper.com.br/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBMenu struct {
	MongoDB.MongoDB
	ClientMenu *mongo.Collection
}

type menuRef struct {
	id   string
	ref  *[]dataFormat.Menu
	pass bool
}

func (e *MongoDBMenu) createTableMainMenu() (err error) {
	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)
	return
}

func (e *MongoDBMenu) GetClassroomMenuFields() (menu []dataFormat.Menu, length int, err error) {
	var cursor *mongo.Cursor

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)
	cursor, err = e.ClientMenu.Find(e.Ctx, bson.M{})
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &menu)
	if err != nil {
		util.TraceToLog()
		return
	}

	length = len(menu)
	return
}

func (e *MongoDBMenu) Get() (menu []dataFormat.Menu, length int, err error) {
	var ref = make([]menuRef, 0)

	menu, err = e.getBySecondaryId("")
	if err != nil {
		return
	}

	for k := range menu {
		ref = append(ref, menuRef{id: menu[k].Id, ref: &menu[k].Menu})
	}

	for {
		err = e.getReference(&ref)
		if err != nil {
			return
		}

		e.clearReference(&ref)

		if len(ref) == 0 {
			break
		}
	}

	length = len(menu)

	return
}

func (e *MongoDBMenu) getReference(ref *[]menuRef) (err error) {
	var menu []dataFormat.Menu

	if ref == nil {
		return
	}

	for refKey := range *ref {
		menu, err = e.getBySecondaryId((*ref)[refKey].id)
		*(*ref)[refKey].ref = menu
		(*ref)[refKey].pass = true
		for menuKey := range menu {
			if menu[menuKey].IdSecondary != "" {
				*ref = append(*ref, menuRef{id: menu[menuKey].Id, ref: &menu[menuKey].Menu})
			}
		}
	}

	return
}

func (e *MongoDBMenu) getBySecondaryId(secondaryId string) (menu []dataFormat.Menu, err error) {
	var cursor *mongo.Cursor
	var menuToQuery = dataFormat.Menu{IdSecondary: secondaryId}

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)
	cursor, err = e.ClientMenu.Find(e.Ctx, menuToQuery.GetIdAndSecondaryIdAsBSonQuery())
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &menu)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

func (e *MongoDBMenu) CreateTableMenu() (err error) {

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)
	return
}

func (e *MongoDBMenu) clearReference(ref *[]menuRef) {
	if ref == nil {
		return
	}

	for {
		pass := false
		for k := range *ref {
			if (*ref)[k].pass == true {
				*ref = append((*ref)[:k], (*ref)[k+1:]...)
				pass = true
				break
			}
		}

		if pass == false {
			break
		}
	}
}

func (e *MongoDBMenu) Install() (err error) {
	var installed = false

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)

	installed, err = e.verifyInstallMenu()
	if err != nil {
		util.TraceToLog()
		return
	}

	if installed == false {
		err = e.CreateTableMenu()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.populateInitialMenu()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *MongoDBMenu) verifyInstallMenu() (installed bool, err error) {
	var cursor *mongo.Cursor
	var users []dataFormat.User
	var userToQuery = dataFormat.User{Id: constants.KMainUserID, Mail: constants.KMainUserMail}

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)

	cursor, err = e.ClientMenu.Find(e.Ctx, userToQuery.GetIdAndMailAsBSonQuery())
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &users)
	if err != nil {
		util.TraceToLog()
		return
	}

	installed = len(users) > 0
	return
}

func (e *MongoDBMenu) Set(id, idSecondary string, typeContent int, classroom int, text string, admin int, icon, url string, order int) (err error) {

	_, err = e.ClientMenu.InsertOne(
		e.Ctx,
		dataFormat.Menu{
			Id:          id,
			IdSecondary: idSecondary,
			Text:        text,
			Admin:       admin,
			Icon:        icon,
			Url:         url,
			ItemOrder:   order,
			TypeContent: typeContent,
			Classroom:   classroom,
			Menu:        nil,
		},
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

func (e *MongoDBMenu) populateInitialMenu() (err error) {

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

func New() (referenceInicialized *MongoDBMenu, err error) {
	referenceInicialized = &MongoDBMenu{}
	err = referenceInicialized.Connect(constants.KMongoDBConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
