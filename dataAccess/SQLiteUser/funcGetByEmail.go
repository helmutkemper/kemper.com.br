package SQLiteUser

import (
	"database/sql"
	"errors"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
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
		return
	}

	var id int
	var menuId int
	var admin int
	var name string
	var nickName string
	var password string

	if rows.Next() {
		err = rows.Scan(&id, &menuId, &admin, &name, &nickName, &mail, &password)
		if err != nil {
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
		err = errors.New("user not found")
	}

	return
}
