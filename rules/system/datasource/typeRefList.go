package datasource

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

// RefList (Português): Recebe todos os ponteiros de datasource
type RefList struct {
	Menu interfaces.InterfaceMenu
	User interfaces.InterfaceUser
}
