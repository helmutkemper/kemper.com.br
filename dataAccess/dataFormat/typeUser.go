package dataFormat

import (
	"encoding/json"
	"github.com/helmutkemper/kemper.com.br/util"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type User struct {
	Id       string `json:"id" bson:"_id"`
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
		Admin    int    `json:"admin"`
		Name     string `json:"name"`
		NickName string `json:"nickname"`
		Mail     string `json:"mail"`
	}{
		Id:       e.Id,
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

func (e *User) GetIdAndMailAsBSonQuery() (query bson.M) {

	var ok bool
	var idElement, mailElement reflect.StructField

	idElement, ok = reflect.ValueOf(*e).Type().FieldByName("Id")
	if ok == false {
		util.TraceToLog()
		return
	}

	mailElement, ok = reflect.ValueOf(*e).Type().FieldByName("Mail")
	if ok == false {
		util.TraceToLog()
		return
	}

	idTagName := idElement.Tag.Get("bson")
	mailTagName := mailElement.Tag.Get("bson")
	query = bson.M{idTagName: e.Id, mailTagName: e.Mail}
	return
}
