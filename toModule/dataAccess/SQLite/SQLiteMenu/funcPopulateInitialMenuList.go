package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"log"
)

// populateInitialMenuList (Português): popula a lista de menu após a instalação
func (e *SQLiteMenu) populateInitialMenuList() (err error) {

	err = e.SetMenuList(constants.KMainMenuID, "MenuMain Principal", 0)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenuList().error: %v", err.Error())
	}

	return
}
