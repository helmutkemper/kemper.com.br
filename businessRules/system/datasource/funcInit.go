package datasource

import (
	"errors"
	"github.com/helmutkemper/kemper.com.br/toModule/passwordHash"
	"github.com/helmutkemper/kemper.com.br/toModule/uID"
	"log"
)

//*************************************************************************
//**                               Cuidado                               **
//*************************************************************************

// Init (Português): Inicializa o datasource escolhido
//   name: tyme Name
//     KSQLite: Inicializa o banco de dados como sendo o SQLite
func (e *RefList) Init(name Name) (err error) {
	err = errors.New("please, inicialize data source first")

	// Inicializa o objeto Password
	e.Password = &passwordHash.Password{}

	// Inicializa o objeto UID
	e.UniqueID = &uID.UID{}

	switch name {
	case KSQLite:

		err = e.initSQLiteMenu()
		if err != nil {
			log.Printf("datasource.Init().error: %v", err.Error())
			return
		}

		err = e.initSQLiteUser()
		if err != nil {
			log.Printf("datasource.Init().error: %v", err.Error())
			return
		}
	}

	return
}
