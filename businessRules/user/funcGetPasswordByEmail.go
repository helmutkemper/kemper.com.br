package user

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	systemDatasource "github.com/helmutkemper/kemper.com.br/rules/system/datasource"
)

func (e *BusinessRules) getPasswordByEmail(mail string) (password string, err error) {
	var userFromDatasource dataFormat.User

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()
	userFromDatasource, err = e.DataSource.GetByEmail(mail)
	if err != nil {
		return
	}

	password = userFromDatasource.Password
	return
}
