package menu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	systemDatasource "github.com/helmutkemper/kemper.com.br/rules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/view/viewMenu"
)

type BusinessRules struct {
	DataSource interfaces.InterfaceMenu `json:"-"`
}

// Menu: (Português): Get menu para o datasource do componente Kendo UI JQuery Menu
func (e *BusinessRules) Get(mainMenuId int) (length int, menu viewMenu.Menu, err error) {
	var menuFromDatasource []dataFormat.Menu

	e.DataSource = systemDatasource.ReferencesList.GetMenu()
	menuFromDatasource, length, err = e.DataSource.Get(mainMenuId)
	if err != nil {
		return
	}

	menu = viewMenu.Menu{}
	menu.Parser(&menuFromDatasource)

	return
}
