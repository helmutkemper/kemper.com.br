package languages

import (
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
)

func (e *BusinessRules) GetAll() (length int, languages []dataFormat.Languages, err error) {
	e.DataSource = systemDatasource.Linker.GetReferenceFromLanguages()
	languages, length, err = e.DataSource.GetAll()

	return
}
