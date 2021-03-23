package dataFormat

import (
	"encoding/json"
)

type User struct {
	Id       string `json:"id"`
	MenuId   string `json:"menuId"`
	Admin    int    `json:"admin"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
	Mail     string `json:"email"`
	Password string `json:"-"`
}

//fixme:arquivo
func (e *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id       string `json:"id"`
		MenuId   string `json:"manuId"`
		Admin    int    `json:"admin"`
		Name     string `json:"name"`
		NickName string `json:"nickname"`
		Mail     string `json:"mail"`
	}{
		Id:       e.Id,
		MenuId:   e.MenuId,
		Admin:    e.Admin,
		Name:     e.Name,
		NickName: e.NickName,
		Mail:     e.Mail,
	})
}
