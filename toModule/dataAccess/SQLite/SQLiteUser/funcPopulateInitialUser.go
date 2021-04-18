package SQLiteUser

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"log"
)

// populateInitialMenu (Português): popula o menu com os primeiros dados após a instalação.
func (e *SQLiteUser) populateInitialUser() (err error) {

	err = e.Set(
		constants.KMainUserID,
		constants.KmainMenuUserAdmin,
		constants.KMainUserName,
		constants.KMainUserNickName,
		constants.KMainUserMail,
		constants.KMainUserPassword,
	)
	if err != nil {
		log.Printf("SQLiteUser.populateInitialUser().error: %v", err.Error())
	}
	return
}
