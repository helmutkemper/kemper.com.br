package SQLiteMenu

import (
	"database/sql"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"log"
)

func (e *SQLiteMenu) GetClassroomMenuFields() (menu []dataFormat.Menu, length int, err error) {
	var rows *sql.Rows

	var id string
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
		       secondaryId,
		       text,
		       admin,
		       icon,
		       url,
		       itemOrder
		FROM
		     menu
		WHERE
				classroom = 1
		ORDER BY
		         itemOrder`,
	)
	if err != nil {
		log.Printf("SQLiteMenu.getBySecondaryId().error: %v", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&id, &idSecondary, &text, &admin, &icon, &url, &itemOrder)
		if err != nil {
			log.Printf("SQLiteMenu.getBySecondaryId().error: %v", err.Error())
			return
		}

		menu = append(
			menu,
			dataFormat.Menu{
				Id:          id,
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

	length = len(menu)
	return
}
