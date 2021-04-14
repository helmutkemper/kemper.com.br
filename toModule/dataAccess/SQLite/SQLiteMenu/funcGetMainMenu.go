package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

// GetMainMenu (Português): Retorna o menu escolhido dentro do formato do datasource
func (e *SQLiteMenu) GetMainMenu(menuId string) (menu []dataFormat.Menu, length int, err error) {
	var ref = make([]menuRef, 0)

	menu, err = e.getBySecondaryId(menuId, "")
	if err != nil {
		return
	}

	for k := range menu {
		ref = append(ref, menuRef{id: menu[k].Id, ref: &menu[k].Menu})
	}

	for {
		err = e.getReference(menuId, &ref)
		if err != nil {
			return
		}

		e.clearReference(&ref)

		if len(ref) == 0 {
			break
		}
	}

	length = len(menu)

	return
}
