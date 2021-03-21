package datasource

import (
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLiteMenu"
)

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
