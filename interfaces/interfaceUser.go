package interfaces

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type InterfaceUser interface {
	GetByEmail(mail string) (user dataFormat.User, err error)
	Set(idMenu, admin int, name, nickName, eMail, password string) (err error)
}
