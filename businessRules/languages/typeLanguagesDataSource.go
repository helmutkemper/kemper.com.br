package languages

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

type BusinessRules struct {
	DataSource interfaces.InterfaceLanguage `json:"-"`
	UniqueID   interfaces.InterfaceUID      `json:"-"`
}
