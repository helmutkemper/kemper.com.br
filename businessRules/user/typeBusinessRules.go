package user

import "github.com/helmutkemper/kemper.com.br/interfaces"

type BusinessRules struct {
	Password   interfaces.InterfacePassword `json:"-"`
	DataSource interfaces.InterfaceUser     `json:"-"`
}
