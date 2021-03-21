package user

import (
	"encoding/base64"
	"github.com/helmutkemper/kemper.com.br/businessRules/passwordHash"
)

func (e *BusinessRules) VerifyPasswordByEmail(mail, password string) (match bool, err error) {
	var pass = passwordHash.Password{}
	var hash []byte
	var passwordFromDatasource string

	passwordFromDatasource, err = e.getPasswordByEmail(mail)
	if err != nil {
		return
	}

	_, err = base64.StdEncoding.Decode([]byte(passwordFromDatasource), hash)
	if err != nil {
		return
	}

	match = pass.CheckHash([]byte(password), hash)
	return
}
