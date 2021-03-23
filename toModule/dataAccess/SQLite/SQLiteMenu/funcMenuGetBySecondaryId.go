package SQLiteMenu

import (
	"database/sql"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"log"
)

// getBySecondaryId (Português): ajuda a montar o menu, navegando do nível mais externo para o mais interno.
func (e *SQLiteMenu) getBySecondaryId(menuId, secondaryId string) (menu []dataFormat.Menu, err error) {
	var rows *sql.Rows

	var id string
	var idMenu string
	var idSecondary string
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
					AND secondaryId = ?
		ORDER BY
		         itemOrder`,
		menuId,
		secondaryId,
	)
	if err != nil {
		log.Printf("SQLiteMenu.getBySecondaryId().error: %v", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&id, &idMenu, &idSecondary, &text, &admin, &icon, &url, &itemOrder)
		if err != nil {
			log.Printf("SQLiteMenu.getBySecondaryId().error: %v", err.Error())
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
