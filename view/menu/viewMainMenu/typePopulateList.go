package viewMainMenu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

// populateList (Português): Relaciona o item do datasource com o item do menu
type populateList struct {
	menuFromDataSource *[]dataFormat.Menu
	menuToPopulate     *[]Item
}
