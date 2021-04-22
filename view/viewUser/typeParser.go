package viewUser

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

func (e *User) Parser(user *dataformat.User) {
	*e = User(*user)
	e.Password = ""
}
