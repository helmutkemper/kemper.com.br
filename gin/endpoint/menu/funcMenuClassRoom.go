package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/businessRules/menu"
	"github.com/helmutkemper/kemper.com.br/view/menu/viewMainMenu"
)

func (e *DataSource) MenuGetListClassRoom(c *gin.Context) {
	var err error
	var menuData viewMainMenu.Menu
	var length int

	menuBusinessRules := menu.BusinessRules{}
	length, menuData, err = menuBusinessRules.GetMenuListClassRoom("5996b891-9d3c-4038-af37-cb07f5f0f72d")

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
