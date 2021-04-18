package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/util"
)

func (e *MongoDBUser) populateInitialUser() (err error) {

	err = e.Set(
		constants.KMainUserID,
		constants.KmainMenuUserAdmin,
		constants.KMainUserName,
		constants.KMainUserNickName,
		constants.KMainUserMail,
		constants.KMainUserPassword,
	)
	if err != nil {
		util.TraceToLog()
		return
	}
	return
}
