package menu

import (
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/view/viewMenu"
)

// Menu: (Português): Get menu para o datasource do componente Kendo UI JQuery Menu
func (e *BusinessRules) Get(mainMenuId int) (length int, menu viewMenu.Menu, err error) {
	var menuFromDatasource []dataFormat.Menu

	e.DataSource = systemDatasource.Linker.GetReferenceFromMenu()
	menuFromDatasource, length, err = e.DataSource.Get(mainMenuId)
	if err != nil {
		return
	}

	menu = viewMenu.Menu{}
	menu.Parser(&menuFromDatasource)

	return
}
