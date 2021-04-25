package interfaces

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

type InterfaceUser interface {
	New() (referenceInicialized interface{}, err error)
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetByEmail(mail string) (user dataformat.User, err error)
	Set(id string, admin int, name, nickName, email, password string) (err error)
	MailExists(mail string) (found bool, err error)
}
