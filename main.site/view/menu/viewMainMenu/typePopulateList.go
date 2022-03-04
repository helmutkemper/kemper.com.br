package viewMainMenu

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

// populateList (Português): Relaciona o item do datasource com o item do menu
type populateList struct {
	menuFromDataSource *[]dataformat.Menu
	menuToPopulate     *[]Item
}
