package user

import "github.com/helmutkemper/kemper.com.br/interfaces"

type BusinessRules struct {
	UniqueID   interfaces.InterfaceUID      `json:"-"`
	Password   interfaces.InterfacePassword `json:"-"`
	DataSource interfaces.InterfaceUser     `json:"-"`
}
