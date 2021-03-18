package dataFormat

import (
	"encoding/base64"
	"encoding/json"
	pass "github.com/helmutkemper/kemper.com.br/password"
)

type User struct {
	id       int
	menuId   int
	admin    int
	name     string
	nickName string
	mail     string
	password string
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
		Id:       e.id,
		MenuId:   e.menuId,
		Admin:    e.admin,
		Name:     e.name,
		NickName: e.nickName,
		Mail:     e.mail,
	})
}

func (e *User) Set(id, meniId, Admin int, name, nickname, mail, password string) (err error) {
	var hash []byte
	p := pass.Password{}
	hash, err = p.MakeHash([]byte(password))
	if err != nil {
		return
	}

	password = base64.StdEncoding.EncodeToString(hash)

	e.id = id
	e.menuId = meniId
	e.admin = Admin
	e.name = name
	e.nickName = nickname
	e.mail = mail
	e.password = password

	return
}
