package datasource

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

func (e *RefList) GetReferenceFromLanguages() (datasource interfaces.InterfaceLanguage) {
	return e.Language
}
