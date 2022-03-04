package datasource

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

// GetReferenceFromUser (Português): Retorna o datasource do usuário
func (e *RefList) GetReferenceFromUser() (datasource interfaces.InterfaceUser) {
	return e.User
}
