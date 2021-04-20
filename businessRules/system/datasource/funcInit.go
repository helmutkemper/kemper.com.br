package datasource

import (
	"errors"
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	"log"
	"plugin"

	//"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDBLanguage"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDBMenu"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDBUser"
	//"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteLanguage"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/toModule/passwordHash"
	"github.com/helmutkemper/kemper.com.br/toModule/uID"
	"github.com/helmutkemper/kemper.com.br/util"
)

// Init (Português): Inicializa o datasource escolhido
//   name: tyme Name
//     KSQLite: Inicializa o banco de dados como sendo o SQLite
func (e *RefList) Init(name Name) (err error) {
	err = errors.New(constants.KErrorInicializeDataSourceFirst)

	// Inicializa o objeto Password
	e.Password = &passwordHash.Password{}

	// Inicializa o objeto UID
	e.UniqueID = &uID.UID{}

	switch name {

	case KPlugin:
		e.Menu, err = SQLiteMenu.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		e.User, err = SQLiteUser.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		var language *plugin.Plugin
		var languaSymbol plugin.Symbol
		var ok bool
		language, err = plugin.Open("./plugin/language.sqlite.so")
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

		err = e.Language.Connect(constants.KSQLiteConnectionString)
		if err != nil {
			log.Printf("SQLiteLanguage.New().Connect().error: %v", err.Error())
			return
		}

		err = e.Language.Install()
		if err != nil {
			log.Printf("SQLiteLanguage.New().Install().error: %v", err.Error())
			return
		}

	case KSQLite:

		e.Menu, err = SQLiteMenu.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		e.User, err = SQLiteUser.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		//e.Language, err = SQLiteLanguage.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

	case KMongoDB:
		e.Menu, err = MongoDBMenu.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		e.User, err = MongoDBUser.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		//e.Language, err = MongoDBLanguage.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}

	case KTotalMadness:
		//SQLite
		e.Menu, err = SQLiteMenu.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		//MongoDB
		e.User, err = MongoDBUser.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		//SQLite
		//e.Language, err = SQLiteLanguage.New()
		//if err != nil {
		//	util.TraceToLog()
		//	return
		//}
	}

	return
}
