package common

import (
	"errors"
	"github.com/helmutkemper/kemper.com.br/dataAccess/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	"log"
)

type dataSourceName int

const (
	// kDataSourceSQLite (Português): Define o datasource como sendo o SQLite3
	kDataSourceSQLite dataSourceName = iota
)

// DataSource (Português): Recebe todos os ponteiros de datasource
type DataSource struct {
	Menu interfaces.InterfaceMenu
}

// Init (Português): Inicializa o datasource escolhido
func (e *DataSource) Init(name dataSourceName) (err error) {
	err = errors.New("please, inicialize data source first")

	switch name {
	case kDataSourceSQLite:
		err = e.initSQLite()
		if err != nil {
			return
		}

		err = e.initMenu()
		if err != nil {
			return
		}
	}

	return
}

// InitMenu (Português): Inicializa o datasource do menu
func (e *DataSource) initMenu() (err error) {
	menu := SQLiteMenu.SQLiteMenu{}
	err = menu.Connect("./db/database.sqlite")
	if err != nil {
		return
	}

	err = menu.Install()
	return
}

// GetMenu (Português): Retorna o datasource do menu
func (e *DataSource) GetMenu() (datasource interfaces.InterfaceMenu) {
	return e.Menu
}

// initSQLite (Português): Inicializa o SQLite
func (e *DataSource) initSQLite() (err error) {
	e.Menu = &SQLiteMenu.SQLiteMenu{}
	err = e.Menu.Connect("./db/database.sqlite")
	return
}

// DataSourceCommon (Português): Arquiva globalmente todos os ponteiros de datasource
var DataSourceCommon DataSource

// init (Português): Inicializa o datasorce escolhido antes da execução do código
func init() {
	var err error
	err = DataSourceCommon.Init(kDataSourceSQLite)
	if err != nil {
		log.Fatalf("Menu datasource initialization error: %v", err.Error())
	}
}
