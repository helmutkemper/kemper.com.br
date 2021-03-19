package datasource

import (
	"errors"
	"github.com/helmutkemper/kemper.com.br/dataAccess/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/dataAccess/SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	"log"
)

type Name int

const (
	// KSQLite (Português): Define o datasource como sendo o SQLite3
	KSQLite Name = iota
)

// RefList (Português): Recebe todos os ponteiros de datasource
type RefList struct {
	Menu interfaces.InterfaceMenu
}

// Init (Português): Inicializa o datasource escolhido
func (e *RefList) Init(name Name) (err error) {
	err = errors.New("please, inicialize data source first")

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

func (e *RefList) initUser() (err error) {
	user := SQLiteUser.SQLiteUser{}
	err = user.Connect("./db/database.sqlite")
	if err != nil {
		return
	}

	err = user.Install()
	return
}

// InitMenu (Português): Inicializa o datasource do menu
func (e *RefList) initMenu() (err error) {
	menu := SQLiteMenu.SQLiteMenu{}
	err = menu.Connect("./db/database.sqlite")
	if err != nil {
		return
	}

	err = menu.Install()
	return
}

// GetMenu (Português): Retorna o datasource do menu
func (e *RefList) GetMenu() (datasource interfaces.InterfaceMenu) {
	return e.Menu
}

// initSQLite (Português): Inicializa o SQLite
func (e *RefList) initSQLite() (err error) {
	e.Menu = &SQLiteMenu.SQLiteMenu{}
	err = e.Menu.Connect("./db/database.sqlite")
	return
}

// DataSourceCommon (Português): Arquiva globalmente todos os ponteiros de datasource
var common RefList

// init (Português): Inicializa o datasorce escolhido antes da execução do código
func init() {
	var err error
	err = common.Init(KSQLite)
	if err != nil {
		log.Fatalf("Menu datasource initialization error: %v", err.Error())
	}
}
