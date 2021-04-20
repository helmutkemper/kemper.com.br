package main

import (
	"database/sql"
	"errors"
	"github.com/helmutkemper/kemper.com.br/constants"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"log"
)

func (e *SQLiteLanguage) GetAll() (languagues []dataFormat.Languages, length int, err error) {
	var rows *sql.Rows

	languagues = make([]dataFormat.Languages, 0)

	rows, err = e.Database.Query(
		`
			SELECT
			    id,
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

	var id string
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
			return
		}

		languagues = append(
			languagues,
			dataFormat.Languages{
				Id:   id,
				Name: name,
			},
		)
	}

	length = len(languagues)

	if length == 0 {
		err = errors.New(constants.KErrorLanguageTableEmpty)
		log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
	}

	return
}
