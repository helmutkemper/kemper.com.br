package viewUser

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

func (e *User) Parser(user *dataFormat.User) {
	*e = User(*user)
	e.Password = ""
}
