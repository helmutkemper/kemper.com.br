package interfaces

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type InterfaceMenu interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	GetMainMenu(menuId string) (menu []dataFormat.Menu, length int, err error)
	Install() (err error)
	SetMainMenu(id, idMenu, idSecondary string, typeContent int, classroom int, text string, admin int, icon, url string, order int) (err error)
	SetMenuList(id, menu string, admin int) (err error)
}
