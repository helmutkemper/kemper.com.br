package viewDataSource

type Item struct {
	Text                    string `json:"text,omitempty"`
	SpriteCssClass          string `json:"spriteCssClass,omitempty"`
	DataSpriteCssClassField string `json:"dataSpriteCssClassField,omitempty"`
	Items                   []Item `json:"items,omitempty"`
	Url                     string `json:"url,omitempty"`
}
