package dataFormat

type Menu struct {
	Id          int    `json:"-"`
	SecondaryId int    `json:"-"`
	Text        string `json:"text"`
	Admin       int    `json:"admin"`
	Icon        string `json:"icon"`
	Url         string `json:"url"`
	ItemOrder   int    `json:"-"`
	SubMenu     []Menu `json:"subMenu,omitempty"`
}
