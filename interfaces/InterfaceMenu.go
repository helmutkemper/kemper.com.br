package interfaces

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type InterfaceMenu interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Get() (menu []dataFormat.Menu, length int, err error)
	GetClassroomMenuFields() (menu []dataFormat.Menu, length int, err error)
	Install() (err error)
	Set(id, idSecondary string, typeContent int, classroom int, text string, admin int, icon, url string, order int) (err error)
}
