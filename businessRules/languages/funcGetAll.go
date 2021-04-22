package languages

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
)

func (e *BusinessRules) GetAll() (length int, languages []dataformat.Languages, err error) {
	e.DataSource = systemDatasource.Linker.GetReferenceFromLanguages()
	languages, length, err = e.DataSource.GetAll()

	return
}
