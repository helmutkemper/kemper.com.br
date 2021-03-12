package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
)

type menuRef struct {
	id   int
	ref  *[]dataFormat.Menu
	pass bool
}
