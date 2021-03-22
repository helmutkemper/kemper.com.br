package datasource

import (
	"errors"
)

//*************************************************************************
//**                               Cuidado                               **
//*************************************************************************

// Init (Português): Inicializa o datasource escolhido
func (e *RefList) Init(name Name) (err error) {
	err = errors.New("please, inicialize data source first")

	//e.Password = &passwordHash.Password{}

	switch name {
	case KSQLite:
		err = e.initSQLite()
		if err != nil {
			return
		}

		err = e.initMenu()
		if err != nil {
			return
		}

		err = e.initUser()
		if err != nil {
			return
		}
	}

	return
}
