package SQLiteMenu

import (
	"database/sql"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

// getTypeContent (Português): Retorna todos os elementos do menu que podem receber conteúdo, em formato datasource.
// Nota: O menu foi pensado para haver dois tipos elemento, 0: recebe outro item de menu. 1: recebe conteúdo clicável.
func (e *SQLiteMenu) getTypeContent(menuId int) (menu []dataFormat.Menu, err error) {
	var rows *sql.Rows

	var id int
	var idMenu int
	var idSecondary int
	var text string
	var admin int
	var icon string
	var url string
	var itemOrder int

	menu = make([]dataFormat.Menu, 0)

	rows, err = e.Database.Query(
		`
		SELECT
		       id,
		       menuId,
		       secondaryId,
		       text,
		       admin,
		       icon,
		       url,
		       itemOrder
		FROM
		     menu
		WHERE
		      menuId = ?
					AND typeContent = ?
		ORDER BY
		         itemOrder`,
		menuId,
		1,
	)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&id, &idMenu, &idSecondary, &text, &admin, &icon, &url, &itemOrder)
		if err != nil {
			return
		}

		menu = append(
			menu,
			dataFormat.Menu{
				Id:          id,
				IdMenu:      idMenu,
				IdSecondary: idSecondary,
				Text:        text,
				Admin:       admin,
				Icon:        icon,
				Url:         url,
				ItemOrder:   itemOrder,
				Menu:        make([]dataFormat.Menu, 0),
			},
		)
	}

	return
}
