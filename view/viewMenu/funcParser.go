package viewMenu

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

func (e *Menu) Parser(menuFromDataSource *[]dataFormat.Menu) {
	menuToPopulate := make([]Item, 0)

	refs := e.getRefs(menuFromDataSource, &menuToPopulate)

	for refKey := range refs {
		items := *(refs[refKey]).menuToPopulate

		for k, v := range *(refs[refKey]).menuFromDataSource {
			items[k].Text = v.Text
			items[k].DataSpriteCssClassField = "k-test-menu"
			items[k].Url = v.Url

			if v.Icon != "" {
				items[k].SpriteCssClass = "menuWith " + v.Icon
			}
		}
	}

	*e = menuToPopulate
}
