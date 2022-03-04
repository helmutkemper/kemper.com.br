package menu

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/view/menu/viewMainMenu"
)

// GetMainMenu (Português): Retorna o menu principal
func (e *BusinessRules) GetMainMenu() (length int, menu viewMainMenu.Menu, err error) {
	var menuFromDatasource []dataformat.Menu

	e.DataSource = systemDatasource.Linker.GetReferenceFromMenu()
	menuFromDatasource, length, err = e.DataSource.Get()
	if err != nil {
		return
	}

	menu = viewMainMenu.Menu{}
	menu.Parser(&menuFromDatasource)

	return
}
