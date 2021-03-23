package SQLiteMenu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

type menuRef struct {
	id   string
	ref  *[]dataFormat.Menu
	pass bool
}
