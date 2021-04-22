package main

import (
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"log"
)

func (e *MongoDBLanguage) Set(id, name string) (err error) {

	_, err = e.ClientLanguage.InsertOne(
		e.Ctx,
		dataFormat.Languages{
			Id:   id,
			Name: name,
		},
	)
	if err != nil {
		log.Printf("MongoDBLanguage.Set().error: %v", err.Error())
		return
	}

	return
}
