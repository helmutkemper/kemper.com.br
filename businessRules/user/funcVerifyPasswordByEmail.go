package user

import (
	"encoding/base64"
	"errors"
)

func (e *BusinessRules) VerifyPasswordByEmail(mail, password string) (match bool, err error) {
	var hash []byte
	var passwordFromDatasource string
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		return
	}

	if matched == false {
		err = errors.New("e-mail must be a valid sintax")
		return
	}

	passwordFromDatasource, err = e.getPasswordByEmail(mail)
	if err != nil {
		return
	}

	_, err = base64.StdEncoding.Decode([]byte(passwordFromDatasource), hash)
	if err != nil {
		return
	}

	match = e.Password.CheckHash([]byte(password), hash)
	return
}
