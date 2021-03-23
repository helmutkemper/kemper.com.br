package SQLiteMenu

import (
	"log"
)

// populateInitialMenuList (Português): popula a lista de menu após a instalação
func (e *SQLiteMenu) populateInitialMenuList() (err error) {

	err = e.SetMenuList("5996b891-9d3c-4038-af37-cb07f5f0f72d", "Menu Principal", 0)
	if err != nil {
		log.Printf("SQLiteMenu.populateInitialMenuList().error: %v", err.Error())
	}

	return
}
