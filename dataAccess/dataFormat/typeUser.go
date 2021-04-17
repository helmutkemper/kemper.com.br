package dataFormat

import (
	"encoding/json"
	"github.com/helmutkemper/kemper.com.br/util"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type User struct {
	Id       string `json:"id" bson:"_id"`
	MenuId   string `json:"menuId" bson:"menuId"`
	Admin    int    `json:"admin"  bson:"admin"`
	Name     string `json:"name"  bson:"name"`
	NickName string `json:"nickname"  bson:"nickname"`
	Mail     string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
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

func (e *User) GetMailAsBSonQuery() (query bson.M) {

	element, ok := reflect.ValueOf(*e).Type().FieldByName("Mail")
	if ok == false {
		util.TraceToLog()
		return
	}

	tagName := element.Tag.Get("bson")
	query = bson.M{tagName: e.Mail}
	return
}
