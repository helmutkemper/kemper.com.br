package datasource

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

// GetUser (Português): Retorna o datasource do usuário
func (e *RefList) GetReferenceFromUser() (datasource interfaces.InterfaceUser) {
	return e.User
}
