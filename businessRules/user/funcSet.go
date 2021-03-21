package user

import (
	"encoding/base64"
	"errors"
	systemDatasource "github.com/helmutkemper/kemper.com.br/rules/system/datasource"
)

func (e *BusinessRules) Set(idMenu, admin int, name, nickName, mail, password string) (err error) {
	var hash, hash64 []byte
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		return
	}

	if matched == false {
		err = errors.New("e-mail must be a valid sintax")
		return
	}

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()

	hash, err = e.Password.MakeHash([]byte(password))
	if err != nil {
		return
	}

	base64.StdEncoding.Encode(hash64, hash)

	err = e.DataSource.Set(idMenu, admin, name, nickName, mail, string(hash64))
	return
}
