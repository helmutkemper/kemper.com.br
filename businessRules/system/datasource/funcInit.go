package datasource

import (
	"errors"
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	"github.com/helmutkemper/kemper.com.br/toModule/passwordHash"
	"github.com/helmutkemper/kemper.com.br/toModule/uID"
	"github.com/helmutkemper/util"
	"plugin"
)

// Init (Português): Inicializa o datasource escolhido
//   name: tyme Name
//     KSQLite: Inicializa o banco de dados como sendo o SQLite
func (e *RefList) Init(name Name) (err error) {
	var ok bool

	var newInterface interface{}

	var menu *plugin.Plugin
	var menuSymbol plugin.Symbol

	var user *plugin.Plugin
	var userSymbol plugin.Symbol

	var language *plugin.Plugin
	var languaSymbol plugin.Symbol

	err = errors.New(constants.KErrorInicializeDataSourceFirst)

	// Inicializa o objeto Password
	e.Password = &passwordHash.Password{}

	// Inicializa o objeto UID
	e.UniqueID = &uID.UID{}

	switch name {

	case KPlugin:
		menu, err = plugin.Open("./plugin/menu.sqlite.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		menuSymbol, err = menu.Lookup("Menu")
		if err != nil {
			util.TraceToLog()
			return
		}

		e.Menu, ok = menuSymbol.(interfaces.InterfaceMenu)
		if ok == false {
			err = errors.New("plugin menu conversion into interface menu has an error")
			util.TraceToLog()
			return
		}

		err = e.Menu.Connect(constants.KSQLiteConnectionString)
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.Menu.Install()
		if err != nil {
			util.TraceToLog()
			return
		}

		user, err = plugin.Open("./plugin/user.sqlite.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		userSymbol, err = user.Lookup("User")
		if err != nil {
			util.TraceToLog()
			return
		}

		e.User, ok = userSymbol.(interfaces.InterfaceUser)
		if ok == false {
			err = errors.New("plugin user conversion into interface user has an error")
			util.TraceToLog()
			return
		}

		err = e.User.Connect(constants.KSQLiteConnectionString)
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.User.Install()
		if err != nil {
			util.TraceToLog()
			return
		}

		language, err = plugin.Open("./plugin/languages.mongodb.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		languaSymbol, err = language.Lookup("Language")
		if err != nil {
			util.TraceToLog()
			return
		}

		e.Language, ok = languaSymbol.(interfaces.InterfaceLanguage)
		if ok == false {
			err = errors.New("plugin language conversion into interface language has an error")
			util.TraceToLog()
			return
		}

		newInterface, err = e.Language.New()
		e.Language = newInterface.(interfaces.InterfaceLanguage)

	case KSQLite:

		//e.Menu, err = SQLiteMenu.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

		//e.User, err = SQLiteUser.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

		//e.Language, err = SQLiteLanguage.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

	case KMongoDB:
		//e.Menu, err = MongoDBMenu.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

		//e.User, err = MongoDBUser.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

		//e.Language, err = MongoDBLanguage.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

	case KTotalMadness:
		//SQLite
		//e.Menu, err = SQLiteMenu.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

		//MongoDB
		//e.User, err = MongoDBUser.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

		//SQLite
		//e.Language, err = SQLiteLanguage.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}
	}

	return
}
