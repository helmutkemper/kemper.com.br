package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/src/github.com/common"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/interfaces"
)

type MenuDataSource struct {
	Restful
	DataSource interfaces.InterfaceMenu `json:"-"`
}

const kMainMenuID = 1

func (e *MenuDataSource) Menu(c *gin.Context) {
	var err error
	var menuData []dataFormat.Menu
	var length int

	e.DataSource = common.DataSourceCommon.GetMenu()
	menuData, length, err = e.DataSource.Get(kMainMenuID)

	e.Meta.Error = []string{}
	if err != nil {
		e.Meta.Success = false
		e.Meta.Error = []string{err.Error()}
	}

	e.Meta.Total = length
	e.Meta.Actual = length
	e.Meta.Success = true
	e.Object = menuData

	c.JSON(200, e)
}
