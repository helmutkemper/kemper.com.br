package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/businessRules/menu"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	"github.com/helmutkemper/kemper.com.br/view/viewMenu"
)

type MenuDataSource struct {
	Restful
	DataSource interfaces.InterfaceMenu `json:"-"`
}

// Menu: (Português): Endpoint menu para o datasource do componente Kendo UI JQuery Menu
func (e *MenuDataSource) Menu(c *gin.Context) {
	var err error
	var menuData viewMenu.Menu
	var length int

	menuBusinessRules := menu.BusinessRules{}
	length, menuData, err = menuBusinessRules.Get(1)

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
	e.Object = menuData

	c.JSON(200, e)
}
