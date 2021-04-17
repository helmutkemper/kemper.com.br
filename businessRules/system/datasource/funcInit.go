package datasource

import (
	"errors"
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDBLanguage"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/MongoDB/MongoDBUser"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteLanguage"
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

		e.Language, err = SQLiteLanguage.New()
		if err != nil {
			util.TraceToLog()
			return
		}

	case KMongoDB:
		e.User, err = MongoDBUser.New()
		if err != nil {
			util.TraceToLog()
			return
		}

		e.Language, err = MongoDBLanguage.New()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}
