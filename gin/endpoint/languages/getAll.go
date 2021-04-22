package languages

import (
	"github.com/gin-gonic/gin"
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/kemper.com.br/businessRules/languages"
	"github.com/helmutkemper/kemper.com.br/gin/endpoint/restful"
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

type DataSource struct {
	restful.Restful
	DataSource interfaces.InterfaceLanguage `json:"-"`
}

func (e *DataSource) GetAll(c *gin.Context) {
	var err error
	var langList []dataformat.Languages
	var length int

	menuBusinessRules := languages.BusinessRules{}
	length, langList, err = menuBusinessRules.GetAll()

	e.Meta.Error = []string{}
	if err != nil {
		e.Meta.Success = false
		e.Meta.Error = []string{err.Error()}
		e.Object = []int{}
		c.JSON(200, e)
		return
	}

	e.Meta.Total = length
	e.Meta.Actual = length
	e.Meta.Success = true
	e.Object = langList

	c.JSON(200, e)
}
