package MongoDBUser

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/util"
)

func (e *MongoDBUser) MailExists(mail string) (found bool, err error) {
	var user dataFormat.User
	user, err = e.GetByEmail(mail)
	if err != nil {
		util.TraceToLog()
		return
	}

	found = user.Mail != ""
	return
}
