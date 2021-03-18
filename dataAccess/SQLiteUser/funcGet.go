package SQLiteUser

import (
	"database/sql"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

// Get (Português): Retorna o menu escolhido dentro do formato do datasource
func (e *SQLiteUser) Get(mail string) (menu []dataFormat.Menu, length int, err error) {
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

	for rows.Next() {
		err = rows.Scan(&id, &menuId, &admin, &name, &nickName, &mail, &password)
		if err != nil {
			return
		}
	}
	return
}
