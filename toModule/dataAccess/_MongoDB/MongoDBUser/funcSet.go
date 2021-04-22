package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/util"
)

func (e *MongoDBUser) Set(id string, admin int, name, nickName, email, password string) (err error) {

	_, err = e.ClientUser.InsertOne(
		e.Ctx,
		dataFormat.User{
			Id:       id,
			Admin:    admin,
			Name:     name,
			NickName: nickName,
			Mail:     email,
			Password: password,
		},
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
