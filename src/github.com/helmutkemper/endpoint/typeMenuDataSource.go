package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/common"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/interfaces"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/viewDataSource"
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

	var m = viewDataSource.Menu{}
	m.CommonDataConvert(&menuData)

	e.Meta.Total = length
	e.Meta.Actual = length
	e.Meta.Success = true
	e.Object = m

	//b, _ := json.Marshal(&e.Object)
	c.JSON(200, e)
	//c.Writer.Write([]byte("callback("))
	//c.Writer.Write(b)
	//c.Writer.Write([]byte(")"))
}
