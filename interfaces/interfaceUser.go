package interfaces

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type InterfaceUser interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetByEmail(mail string) (user dataFormat.User, err error)
	Set(id, idMenu string, admin int, name, nickName, eMail, password string) (err error)
}
