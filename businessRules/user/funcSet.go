package user

import (
	"encoding/base64"
	"github.com/helmutkemper/kemper.com.br/businessRules/passwordHash"
	systemDatasource "github.com/helmutkemper/kemper.com.br/rules/system/datasource"
)

func (e *BusinessRules) Set(idMenu, admin int, name, nickName, eMail, password string) (err error) {
	var pass = passwordHash.Password{}
	var hash, hash64 []byte

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()

	hash, err = pass.MakeHash([]byte(password))
	if err != nil {
		return
	}

	base64.StdEncoding.Encode(hash64, hash)

	err = e.DataSource.Set(idMenu, admin, name, nickName, eMail, string(hash64))
	return
}
