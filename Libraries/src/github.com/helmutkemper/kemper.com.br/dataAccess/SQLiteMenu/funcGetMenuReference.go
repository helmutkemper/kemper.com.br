package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

// getReference (Português): Gera ponteiros para todos os submenus, permitindo a navegação durante o prenchimento de
// dados.
func (e *SQLiteMenu) getReference(menuId int, ref *[]menuRef) (err error) {
	var menu []dataFormat.Menu

	if ref == nil {
		return
	}

	for refKey := range *ref {
		menu, err = e.getBySecondaryId(menuId, (*ref)[refKey].id)
		*(*ref)[refKey].ref = menu
		(*ref)[refKey].pass = true
		for menuKey := range menu {
			if menu[menuKey].IdSecondary != 0 {
				*ref = append(*ref, menuRef{id: menu[menuKey].Id, ref: &menu[menuKey].Menu})
			}
		}
	}

	return
}
