package interfaces

import (
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
)

type InterfaceMenu interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Get(menuId int) (menu []dataFormat.Menu, err error)
	Install() (err error)
	SetMenu(idMenu, idSecondary int, text string, admin int, icon, url string, order int) (err error)
}
