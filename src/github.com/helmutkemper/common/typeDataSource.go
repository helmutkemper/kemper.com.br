package common

import (
	"errors"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/interfaces"
	"log"
)

type dataSourceName int

const (
	kDataSourceSQLite dataSourceName = iota
)

type DataSource struct {
	Menu interfaces.InterfaceMenu
}

func (e *DataSource) Init(name dataSourceName) (err error) {
	err = errors.New("please, inicialize data source first")

	if name == kDataSourceSQLite {
		err = e.InitSQLite()
		if err != nil {
			return
		}
	}

	return
}

func (e *DataSource) GetMenu() (datasource interfaces.InterfaceMenu) {
	return e.Menu
}

func (e *DataSource) InitSQLite() (err error) {
	e.Menu = &SQLiteMenu.SQLiteMenu{}
	err = e.Menu.Connect("./db/database.sqlite")
	return
}

var DataSourceCommon DataSource

func init() {
	var err error
	err = DataSourceCommon.Init(kDataSourceSQLite)
	if err != nil {
		log.Fatalf("Menu datasource initialization error: %v", err.Error())
	}
}
