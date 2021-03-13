package dataSource

type Item struct {
	Id                      int    `json:"-"`
	IdSecondary             int    `json:"-"`
	Text                    string `json:"text,omitempty"`
	SpriteCssClass          string `json:"spriteCssClass,omitempty"`
	DataSpriteCssClassField string `json:"dataSpriteCssClassField,omitempty"`
	Items                   []Item `json:"items,omitempty"`
	Url                     string `json:"url,omitempty"`
}
