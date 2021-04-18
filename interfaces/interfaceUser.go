package interfaces

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type InterfaceUser interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetByEmail(mail string) (user dataFormat.User, err error)
	Set(id string, admin int, name, nickName, email, password string) (err error)
	MailExists(mail string) (found bool, err error)
}
