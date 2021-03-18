package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type menuRef struct {
	id   int
	ref  *[]dataFormat.Menu
	pass bool
}
