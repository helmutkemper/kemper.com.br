package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
)

// Get (Português): Retorna o menu escolhido dentro do formato do datasource
func (e *SQLiteMenu) Get(menuId int) (menu []dataFormat.Menu, length int, err error) {
	var ref = make([]menuRef, 0)

	menu, err = e.getBySecondaryId(menuId, 0)
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
