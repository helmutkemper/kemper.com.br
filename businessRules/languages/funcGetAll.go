package languages

import (
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
)

func (e *BusinessRules) GetAll() (length int, languages []string, err error) {
	e.DataSource = systemDatasource.Linker.GetReferenceFromLanguages()
	languages, length, err = e.DataSource.GetAll()

	return
}
