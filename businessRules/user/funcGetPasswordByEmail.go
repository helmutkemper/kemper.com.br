package user

import (
	"errors"
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"log"
)

func (e *BusinessRules) getPasswordByEmail(mail string) (password string, err error) {
	var userFromDatasource dataFormat.User
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New("e-mail must be a valid sintax")
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()
	userFromDatasource, err = e.DataSource.GetByEmail(mail)
	if err != nil {
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	password = userFromDatasource.Password
	return
}
