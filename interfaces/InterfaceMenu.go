package interfaces

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

type InterfaceMenu interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Get() (menu []dataformat.Menu, length int, err error)
	GetClassroomMenuFields() (menu []dataformat.Menu, length int, err error)
	Install() (err error)
	Set(id, idSecondary string, typeContent int, classroom int, text string, admin int, icon, url string, order int) (err error)
}
