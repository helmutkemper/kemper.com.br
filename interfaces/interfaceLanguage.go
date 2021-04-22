package interfaces

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

type InterfaceLanguage interface {
	New() (referenceInicialized interface{}, err error)
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetAll() (languagues []dataformat.Languages, length int, err error)
	Set(id, name string) (err error)
}
