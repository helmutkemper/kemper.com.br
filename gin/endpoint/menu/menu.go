package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/businessRules/menu"
	"github.com/helmutkemper/kemper.com.br/gin/endpoint/restful"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	"github.com/helmutkemper/kemper.com.br/view/menu/viewMainMenu"
)

type DataSource struct {
	restful.Restful
	DataSource interfaces.InterfaceMenu `json:"-"`
}

// Menu (Português): Endpoint menu para o datasource do componente Kendo UI JQuery Menu
func (e *DataSource) Menu(c *gin.Context) {
	var err error
	var menuData viewMainMenu.Menu
	var length int

	menuBusinessRules := menu.BusinessRules{}
	length, menuData, err = menuBusinessRules.GetMainMenu("5996b891-9d3c-4038-af37-cb07f5f0f72d")

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
