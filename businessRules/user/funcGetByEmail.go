package user

import (
	"errors"
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/view/viewUser"
	"log"
)

// Get: (Português):
func (e *BusinessRules) GetByEmail(mail string) (length int, user viewUser.User, err error) {
	var userFromDatasource dataFormat.User
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.GetByEmail().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New("e-mail must be a valid sintax")
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
