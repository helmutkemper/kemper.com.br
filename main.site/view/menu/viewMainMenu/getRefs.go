package viewMainMenu

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

// getRefs (Português): Gera uma lista de referências dos itens do datasource para ser usado na função Parser()
func (e *Menu) getRefs(menuFromDataSource *[]dataformat.Menu, menuToPopulate *[]Item) (pointer []populateList) {
	pointer = make([]populateList, 0)
	*menuToPopulate = make([]Item, len(*menuFromDataSource))

	pointer = append(
		pointer,
		populateList{
			menuFromDataSource: menuFromDataSource,
			menuToPopulate:     menuToPopulate,
		},
	)
	for k := range *menuFromDataSource {
		if (*menuFromDataSource)[k].Menu != nil {
			pointer = append(
				pointer,
				populateList{
					menuFromDataSource: &(*menuFromDataSource)[k].Menu,
					menuToPopulate:     &(*menuToPopulate)[k].Items,
				},
			)
			pointer = append(pointer, e.getRefs(&(*menuFromDataSource)[k].Menu, &(*menuToPopulate)[k].Items)...)
		}
	}

	return
}
