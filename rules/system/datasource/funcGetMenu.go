package datasource

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

// GetMenu (Português): Retorna o datasource do menu
func (e *RefList) GetMenu() (datasource interfaces.InterfaceMenu) {
	return e.Menu
}