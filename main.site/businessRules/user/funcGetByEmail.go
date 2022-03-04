package user

import (
	"errors"
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/view/viewUser"
	"log"
)

// GetByEmail (Português):
func (e *BusinessRules) GetByEmail(mail string) (length int, user viewUser.User, err error) {
	var userFromDatasource dataformat.User
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.GetByEmail().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New(constants.KErrorEmailValidSintax)
		log.Printf("user.GetByEmail().error: %v", err.Error())
		return
	}

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()
	userFromDatasource, err = e.DataSource.GetByEmail(mail)
	if err != nil {
		log.Printf("user.GetByEmail().error: %v", err.Error())
		return
	}

	if userFromDatasource.Mail != "" {
		length = 1
	}

	user = viewUser.User{}
	user.Parser(&userFromDatasource)

	return
}
