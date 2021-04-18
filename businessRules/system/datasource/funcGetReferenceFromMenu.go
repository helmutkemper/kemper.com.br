package datasource

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

// GetReferenceFromMenu (Português): Retorna o datasource do menu
func (e *RefList) GetReferenceFromMenu() (datasource interfaces.InterfaceMenu) {
	return e.Menu
}
