package dataFormat

import (
	"encoding/json"
)

type User struct {
	Id       int
	MenuId   int
	Admin    int
	Name     string
	NickName string
	Mail     string
	Password string
}

func (e *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id       int    `json:"id"`
		MenuId   int    `json:"manuId"`
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
