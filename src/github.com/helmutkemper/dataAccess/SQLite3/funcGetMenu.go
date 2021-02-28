package sqlite3

import (
	"database/sql"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
)

func (e *SQLite3) GetMenu(subMenuId int, subMenu *[]dataFormat.Menu) (mountedMenu []dataFormat.Menu, err error) {
	var rows *sql.Rows

	var id int
	var text string
	var admin int
	var icon string
	var url string

	var menu = &[]dataFormat.Menu{}

	if subMenu != nil {
		menu = subMenu
	}

	rows, err = e.database.Query(
		`SELECT id, text, admin, icon, url FROM menu WHERE secondaryId = ? ORDER BY itemOrder`,
		subMenuId,
	)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&id, &text, &admin, &icon, &url)
		if err != nil {
			return
		}

		var m = dataFormat.Menu{
			Text:  text,
			Admin: admin,
			Icon:  icon,
			Url:   url,
		}

		_, err = e.GetMenu(id, &m.SubMenu)
		if err != nil {
			return
		}

		*menu = append(*menu, m)
	}

	mountedMenu = *menu

	return
}
