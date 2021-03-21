package user

import (
	"errors"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	systemDatasource "github.com/helmutkemper/kemper.com.br/rules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/view/viewUser"
)

// Get: (Português):
func (e *BusinessRules) GetByEmail(mail string) (length int, user viewUser.User, err error) {
	var userFromDatasource dataFormat.User
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
	userFromDatasource, err = e.DataSource.GetByEmail(mail)
	if err != nil {
		return
	}

	if user.Mail != "" {
		length = 1
	}

	user = viewUser.User{}
	user.Parser(&userFromDatasource)

	return
}
