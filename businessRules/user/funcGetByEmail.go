package user

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	systemDatasource "github.com/helmutkemper/kemper.com.br/rules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/view/viewUser"
)

// Get: (Português):
func (e *BusinessRules) GetByEmail(mail string) (length int, user viewUser.User, err error) {
	var userFromDatasource dataFormat.User

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
