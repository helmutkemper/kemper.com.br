package main

import (
	"github.com/helmutkemper/kemper.com.br/constants"
	"log"
)

func (e *SQLiteLanguage) New() (referenceInicialized *SQLiteLanguage, err error) {
	referenceInicialized = &SQLiteLanguage{}
	err = referenceInicialized.Connect(constants.KSQLiteConnectionString)
	if err != nil {
		log.Printf("SQLiteLanguage.New().Connect().error: %v", err.Error())
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		log.Printf("SQLiteLanguage.New().Install().error: %v", err.Error())
		return
	}

	return
}
