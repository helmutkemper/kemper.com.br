package viewMenu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type populateList struct {
	menuFromDataSource *[]dataFormat.Menu
	menuToPopulate     *[]Item
}
