package dataFormat

type Menu struct {
	Id          string `json:"id"`
	IdMenu      string `json:"menuId"`
	IdSecondary string `json:"secondaryId"`
	Text        string `json:"text"`
	Admin       int    `json:"admin"`
	Icon        string `json:"icon,omitempty"`
	Url         string `json:"url,omitempty"`
	ItemOrder   int    `json:"itemOrder"`
	Menu        []Menu `json:"menu,omitempty"`
}
