package datasource

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

// RefList (Português): Recebe todos os ponteiros de datasource
type RefList struct {
	Menu     interfaces.InterfaceMenu     `json:"-"`
	User     interfaces.InterfaceUser     `json:"-"`
	Language interfaces.InterfaceLanguage `json:"-"`
	Password interfaces.InterfacePassword `json:"-"`
	UniqueID interfaces.InterfaceUID      `json:"-"`
	Jwt      interfaces.InterfaceJWT      `json:"-"`
}
