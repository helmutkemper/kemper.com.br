package businessRules

import "github.com/helmutkemper/kemper.com.br/interfaces"

type BusinessRules struct {
	UniqueKey interfaces.UniqueKey
	Password  interfaces.Password
	Data      interfaces.DataAccess
}
