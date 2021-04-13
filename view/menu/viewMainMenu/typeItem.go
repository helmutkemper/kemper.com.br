package viewMainMenu

// Item (Português): Item do menu a ser usado pelo menu do kendo ui
type Item struct {
	Text                    string `json:"text,omitempty"`
	SpriteCssClass          string `json:"spriteCssClass,omitempty"`
	DataSpriteCssClassField string `json:"dataSpriteCssClassField,omitempty"`
	Items                   []Item `json:"items,omitempty"`
	Url                     string `json:"url,omitempty"`
}
