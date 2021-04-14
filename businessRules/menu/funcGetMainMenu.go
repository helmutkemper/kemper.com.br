package menu

import (
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/view/menu/viewMainMenu"
)

// GetMainMenu (Português): Retorna o menu principal
func (e *BusinessRules) GetMainMenu(mainMenuId string) (length int, menu viewMainMenu.Menu, err error) {
	var menuFromDatasource []dataFormat.Menu

	e.DataSource = systemDatasource.Linker.GetReferenceFromMenu()
	menuFromDatasource, length, err = e.DataSource.GetMenuMain(mainMenuId)
	if err != nil {
		return
	}

	menu = viewMainMenu.Menu{}
	menu.Parser(&menuFromDatasource)

	return
}
