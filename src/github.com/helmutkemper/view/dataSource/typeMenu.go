package dataSource

import (
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
)

type Item struct {
	Text                    string `json:"text,omitempty"`
	SpriteCssClass          string `json:"spriteCssClass,omitempty"`
	DataSpriteCssClassField string `json:"dataSpriteCssClassField,omitempty"`
	Items                   []Item `json:"items,omitempty"`
	Url                     string `json:"url,omitempty"`
}

type Menu []Item

func (e *Menu) CommonDataConvert(menu *[]dataFormat.Menu, reference *[]Item) {
	var m []Item
	if e == nil {
		*e = make([]Item, 0)
	}

	for _, item := range *menu {
		toDataSourceItem := Item{}
		toDataSourceItem.Text = item.Text
		toDataSourceItem.SpriteCssClass = "menuWith " + item.Icon
		toDataSourceItem.DataSpriteCssClassField = "k-test-menu"
		toDataSourceItem.Url = item.Url

		if reference == nil {
			*e = append(*e, toDataSourceItem)
		} else {
			*reference = append(*reference, toDataSourceItem)
		}

		if item.SubMenu != nil {
			if reference == nil {
				reference = &m
			}

			e.CommonDataConvert(&item.SubMenu, reference)

			//(*e)[k].Items = *reference
		}
	}
}
