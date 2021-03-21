package user

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

type BusinessRules struct {
	DataSource interfaces.InterfaceUser `json:"-"`
}
