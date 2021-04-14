package SQLiteLanguage

import (
	"database/sql"
	"errors"
	"github.com/helmutkemper/kemper.com.br/constants"
	"log"
)

func (e *SQLiteLanguage) GetAll() (languagues []string, lenght int, err error) {
	var rows *sql.Rows

	languagues = make([]string, 0)

	rows, err = e.Database.Query(
		`
			SELECT
				name
			FROM
				language
			ORDER BY 
				name
		`,
	)
	if err != nil {
		log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
		return
	}

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
			return
		}

		languagues = append(languagues, name)
	}

	lenght = len(languagues)

	if lenght == 0 {
		err = errors.New(constants.KErrorLanguageTableEmpty)
		log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
	}

	return
}
