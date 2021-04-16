package interfaces

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type InterfaceLanguage interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetAll() (languagues []dataFormat.Languages, lenght int, err error)
	Set(id, name string) (err error)
}
