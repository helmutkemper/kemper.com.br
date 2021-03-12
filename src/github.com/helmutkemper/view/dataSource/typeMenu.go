package dataSource

import (
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
)

type Item struct {
	Id                      int    `json:"-"`
	IdSecondary             int    `json:"-"`
	Text                    string `json:"text,omitempty"`
	SpriteCssClass          string `json:"spriteCssClass,omitempty"`
	DataSpriteCssClassField string `json:"dataSpriteCssClassField,omitempty"`
	Items                   []Item `json:"items,omitempty"`
	Url                     string `json:"url,omitempty"`
}

type Menu []Item

func (e *Menu) CommonDataConvert(menu *[]dataFormat.Menu) {
	*e = make([]Item, 0)

	for _, v := range *menu {
		if v.IdSecondary == 0 {
			item := Item{
				Id:                      v.Id,
				IdSecondary:             v.IdSecondary,
				Text:                    v.Text,
				SpriteCssClass:          "",
				DataSpriteCssClassField: "",
				Items:                   make([]Item, 0),
				Url:                     v.Url,
			}
			*e = append(*e, item)
		}
	}
}

func (e *Menu) commonDataConvert(menu *[]dataFormat.Menu, reference *[]Item) {

}
