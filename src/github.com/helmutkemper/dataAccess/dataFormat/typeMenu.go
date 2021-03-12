package dataFormat

type Menu struct {
	Id          int    `json:"id"`
	IdMenu      int    `json:"menuId"`
	IdSecondary int    `json:"secondaryId"`
	Text        string `json:"text"`
	Admin       int    `json:"admin"`
	Icon        string `json:"icon"`
	Url         string `json:"url"`
	ItemOrder   int    `json:"itemOrder"`
	Menu        []Menu `json:"menu"`
}
