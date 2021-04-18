package dataFormat

type MenuList struct {
	Id    string `json:"id" bson:"_id"`
	Menu  string `json:"menu" bson:"menu"`
	Admin int    `json:"admin" bson:"admin"`
}
