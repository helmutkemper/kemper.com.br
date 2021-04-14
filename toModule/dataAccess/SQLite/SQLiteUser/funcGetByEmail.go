package SQLiteUser

import (
	"database/sql"
	"errors"
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"log"
)

// GetByEmail (Português): Retorna o menu escolhido dentro do formato do datasource
func (e *SQLiteUser) GetByEmail(mail string) (user dataFormat.User, err error) {
	var rows *sql.Rows
	rows, err = e.Database.Query(
		`
			SELECT
				id,
				menuId,
				admin,
				name,
				nickName,
				eMail,
				password
			FROM
				user
			WHERE
				eMail = ?
		`,
		mail,
	)
	if err != nil {
		log.Printf("SQLiteUser.GetByEmail().error: %v", err.Error())
		return
	}

	var id string
	var menuId string
	var admin int
	var name string
	var nickName string
	var password string

	if rows.Next() {
		err = rows.Scan(&id, &menuId, &admin, &name, &nickName, &mail, &password)
		if err != nil {
			log.Printf("SQLiteUser.GetByEmail().error: %v", err.Error())
			return
		}

		user.Id = id
		user.MenuId = menuId
		user.Admin = admin
		user.Name = name
		user.NickName = nickName
		user.Mail = mail
		user.Password = password
	} else {
		err = errors.New(constants.KErrorUserNotFound)
		log.Printf("SQLiteUser.GetByEmail().error: %v", err.Error())
	}

	return
}
